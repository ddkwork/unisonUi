/*
 * Copyright ©1998-2022 by Richard A. Wilkes. All rights reserved.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, version 2.0. If a copy of the MPL was not distributed with
 * this file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * This Source Code Form is "Incompatible With Secondary Licenses", as
 * defined by the Mozilla Public License, version 2.0.
 */

package ux

import (
	"fmt"

	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/gcs/v5/svg"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/unison"
)

// PointsPanel holds the contents of the points block on the sheet.
type PointsPanel struct {
	unison.Panel
	entity       *model.Entity
	targetMgr    *TargetMgr
	prefix       string
	total        *unison.Label
	ptsList      *unison.Panel
	unspentLabel *unison.Label
	overSpent    int8
}

// NewPointsPanel creates a new points panel.
func NewPointsPanel(entity *model.Entity, targetMgr *TargetMgr) *PointsPanel {
	p := &PointsPanel{
		entity:    entity,
		targetMgr: targetMgr,
		prefix:    targetMgr.NextPrefix(),
	}
	p.Self = p
	p.SetLayout(&unison.FlexLayout{Columns: 1})
	p.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.EndAlignment,
		VAlign: unison.FillAlignment,
		VSpan:  2,
		VGrab:  true,
	})

	hdr := unison.NewPanel()
	hdr.SetLayout(&unison.FlexLayout{
		Columns: 1,
		HAlign:  unison.MiddleAlignment,
	})
	hdr.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})
	hdr.DrawCallback = func(gc *unison.Canvas, rect unison.Rect) {
		gc.DrawRect(rect, model.HeaderColor.Paint(gc, rect, unison.Fill))
	}

	hdri := unison.NewPanel()
	hdri.SetLayout(&unison.FlexLayout{
		Columns:  2,
		HSpacing: 4,
	})
	hdri.SetLayoutData(&unison.FlexLayoutData{HAlign: unison.MiddleAlignment})
	hdr.AddChild(hdri)

	var overallTotal string
	if p.entity.SheetSettings.ExcludeUnspentPointsFromTotal {
		overallTotal = p.entity.SpentPoints().String()
	} else {
		overallTotal = p.entity.TotalPoints.String()
	}
	p.total = unison.NewLabel()
	p.total.Font = model.PageLabelPrimaryFont
	p.total.Text = fmt.Sprintf(i18n.Text("%s Points"), overallTotal)
	p.total.OnBackgroundInk = model.OnHeaderColor
	hdri.AddChild(p.total)
	height := p.total.Font.Baseline() - 2
	editButton := unison.NewSVGButton(svg.Edit)
	editButton.Font = model.PageLabelPrimaryFont
	editButton.Drawable.(*unison.DrawableSVG).Size = unison.NewSize(height, height)
	editButton.ClickCallback = func() {
		displayPointsEditor(unison.AncestorOrSelf[Rebuildable](p), p.entity)
	}
	hdri.AddChild(editButton)
	p.AddChild(hdr)

	p.ptsList = unison.NewPanel()
	p.ptsList.SetLayout(&unison.FlexLayout{
		Columns:  2,
		HSpacing: 4,
	})
	p.ptsList.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.EndAlignment,
		VAlign: unison.FillAlignment,
		VSpan:  2,
		VGrab:  true,
	})
	p.AddChild(p.ptsList)

	p.ptsList.SetBorder(unison.NewCompoundBorder(unison.NewLineBorder(model.HeaderColor, 0, unison.Insets{
		Top:    0,
		Left:   1,
		Bottom: 1,
		Right:  1,
	}, false), unison.NewEmptyBorder(unison.Insets{
		Top:    1,
		Left:   2,
		Bottom: 1,
		Right:  2,
	})))
	p.ptsList.DrawCallback = func(gc *unison.Canvas, rect unison.Rect) { drawBandedBackground(p.ptsList, gc, rect, 0, 2) }

	p.unspentLabel = p.addPointsField(NewNonEditablePageFieldEnd(func(f *NonEditablePageField) {
		if text := p.entity.UnspentPoints().String(); text != f.Text {
			f.Text = text
			p.adjustUnspent()
			MarkForLayoutWithinDockable(f)
		}
	}), i18n.Text("Unspent"), i18n.Text("Points earned but not yet spent"))
	p.unspentLabel.DrawCallback = func(gc *unison.Canvas, rect unison.Rect) {
		if p.overSpent == -1 {
			gc.DrawRect(rect, unison.ErrorColor.Paint(gc, rect, unison.Fill))
		}
		p.unspentLabel.DefaultDraw(gc, rect)
	}
	p.addPointsField(NewNonEditablePageFieldEnd(func(f *NonEditablePageField) {
		_, _, race, _ := p.entity.TraitPoints()
		if text := race.String(); text != f.Text {
			f.Text = text
			MarkForLayoutWithinDockable(f)
		}
	}), i18n.Text("Race"), i18n.Text("Total points spent on a racial package"))
	p.addPointsField(NewNonEditablePageFieldEnd(func(f *NonEditablePageField) {
		if text := p.entity.AttributePoints().String(); text != f.Text {
			f.Text = text
			MarkForLayoutWithinDockable(f)
		}
	}), i18n.Text("Attributes"), i18n.Text("Total points spent on attributes"))
	p.addPointsField(NewNonEditablePageFieldEnd(func(f *NonEditablePageField) {
		ad, _, _, _ := p.entity.TraitPoints()
		if text := ad.String(); text != f.Text {
			f.Text = text
			MarkForLayoutWithinDockable(f)
		}
	}), i18n.Text("Advantages"), i18n.Text("Total points spent on advantages"))
	p.addPointsField(NewNonEditablePageFieldEnd(func(f *NonEditablePageField) {
		_, disad, _, _ := p.entity.TraitPoints()
		if text := disad.String(); text != f.Text {
			f.Text = text
			MarkForLayoutWithinDockable(f)
		}
	}), i18n.Text("Disadvantages"), i18n.Text("Total points spent on disadvantages"))
	p.addPointsField(NewNonEditablePageFieldEnd(func(f *NonEditablePageField) {
		_, _, _, quirk := p.entity.TraitPoints()
		if text := quirk.String(); text != f.Text {
			f.Text = text
			MarkForLayoutWithinDockable(f)
		}
	}), i18n.Text("Quirks"), i18n.Text("Total points spent on quirks"))
	p.addPointsField(NewNonEditablePageFieldEnd(func(f *NonEditablePageField) {
		if text := p.entity.SkillPoints().String(); text != f.Text {
			f.Text = text
			MarkForLayoutWithinDockable(f)
		}
	}), i18n.Text("Skills"), i18n.Text("Total points spent on skills"))
	p.addPointsField(NewNonEditablePageFieldEnd(func(f *NonEditablePageField) {
		if text := p.entity.SpellPoints().String(); text != f.Text {
			f.Text = text
			MarkForLayoutWithinDockable(f)
		}
	}), i18n.Text("Spells"), i18n.Text("Total points spent on spells"))
	p.adjustUnspent()
	return p
}

func (p *PointsPanel) addPointsField(field *NonEditablePageField, title, tooltip string) *unison.Label {
	field.Tooltip = unison.NewTooltipWithText(tooltip)
	p.ptsList.AddChild(field)
	label := NewPageLabel(title)
	label.Tooltip = unison.NewTooltipWithText(tooltip)
	p.ptsList.AddChild(label)
	return label
}

func (p *PointsPanel) adjustUnspent() {
	if p.unspentLabel != nil {
		last := p.overSpent
		if p.entity.UnspentPoints() < 0 {
			if p.overSpent != -1 {
				p.overSpent = -1
				p.unspentLabel.OnBackgroundInk = unison.OnErrorColor
				p.unspentLabel.Text = i18n.Text("Overspent")
			}
		} else {
			if p.overSpent != 1 {
				p.overSpent = 1
				p.unspentLabel.OnBackgroundInk = unison.OnContentColor
				p.unspentLabel.Text = i18n.Text("Unspent")
			}
		}
		if last != p.overSpent {
			MarkForLayoutWithinDockable(p)
		}
	}
}

// Sync the panel to the current data.
func (p *PointsPanel) Sync() {
	var overallTotal string
	if p.entity.SheetSettings.ExcludeUnspentPointsFromTotal {
		overallTotal = p.entity.SpentPoints().String()
	} else {
		overallTotal = p.entity.TotalPoints.String()
	}
	p.total.Text = fmt.Sprintf(i18n.Text("%s Points"), overallTotal)
	p.MarkForLayoutAndRedraw()
}
