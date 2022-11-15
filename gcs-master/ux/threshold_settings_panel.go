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
	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/gcs/v5/svg"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/unison"
)

type thresholdSettingsPanel struct {
	unison.Panel
	pool         *poolSettingsPanel
	threshold    *model.PoolThreshold
	deleteButton *unison.Button
}

func newThresholdSettingsPanel(pool *poolSettingsPanel, threshold *model.PoolThreshold) *thresholdSettingsPanel {
	p := &thresholdSettingsPanel{
		pool:      pool,
		threshold: threshold,
	}
	p.Self = p
	p.SetBorder(unison.NewEmptyBorder(unison.Insets{
		Top:    unison.StdVSpacing,
		Left:   unison.StdHSpacing,
		Bottom: unison.StdVSpacing,
		Right:  unison.StdHSpacing,
	}))
	p.DrawCallback = func(gc *unison.Canvas, rect unison.Rect) {
		color := unison.ContentColor
		if p.Parent().IndexOfChild(p)%2 == 1 {
			color = unison.BandingColor
		}
		gc.DrawRect(rect, color.Paint(gc, rect, unison.Fill))
	}
	p.SetLayout(&unison.FlexLayout{
		Columns:  3,
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	p.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})

	p.AddChild(NewDragHandle(map[string]any{
		attributeSettingsDragDataKey: &attributeSettingsDragData{
			owner:     pool.dockable.Entity(),
			def:       pool.def,
			threshold: threshold,
		},
	}))
	p.AddChild(p.createButtons())
	p.AddChild(p.createContent())
	return p
}

func (p *thresholdSettingsPanel) createButtons() *unison.Panel {
	buttons := unison.NewPanel()
	buttons.SetLayout(&unison.FlexLayout{
		Columns:  1,
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	buttons.SetLayoutData(&unison.FlexLayoutData{HAlign: unison.MiddleAlignment})

	p.deleteButton = unison.NewSVGButton(svg.Trash)
	p.deleteButton.ClickCallback = func() { p.pool.deleteThreshold(p) }
	p.deleteButton.Tooltip = unison.NewTooltipWithText(i18n.Text("Remove pool threshold"))
	p.deleteButton.SetEnabled(len(p.pool.def.Thresholds) > 1)
	buttons.AddChild(p.deleteButton)
	return buttons
}

func (p *thresholdSettingsPanel) createContent() *unison.Panel {
	content := unison.NewPanel()
	content.SetLayout(&unison.FlexLayout{
		Columns:  2,
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	content.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})

	text := i18n.Text("State")
	content.AddChild(NewFieldLeadingLabel(text))
	field := NewStringField(p.pool.dockable.targetMgr, p.threshold.KeyPrefix+"state", text,
		func() string { return p.threshold.State },
		func(s string) { p.threshold.State = s })
	field.SetMinimumTextWidthUsing(prototypeMinIDWidth)
	field.Tooltip = unison.NewTooltipWithText(i18n.Text("A short description of the threshold state"))
	field.SetLayoutData(&unison.FlexLayoutData{HAlign: unison.FillAlignment})
	content.AddChild(field)

	text = i18n.Text("Threshold")
	content.AddChild(NewFieldLeadingLabel(text))
	field = NewStringField(p.pool.dockable.targetMgr, p.threshold.KeyPrefix+"threshold", text,
		func() string { return p.threshold.Expression },
		func(s string) { p.threshold.Expression = s })
	field.SetMinimumTextWidthUsing("round($self*100/50+20)")
	field.Tooltip = unison.NewTooltipWithText(i18n.Text("An expression to calculate the threshold value"))
	content.AddChild(field)

	for _, op := range model.AllThresholdOp[1:] {
		content.AddChild(unison.NewPanel())
		content.AddChild(p.createOpCheckBox(op))
	}

	text = i18n.Text("Explanation")
	content.AddChild(NewFieldLeadingLabel(text))
	field = NewMultiLineStringField(p.pool.dockable.targetMgr, p.threshold.KeyPrefix+"explanation", text,
		func() string { return p.threshold.Explanation },
		func(s string) { p.threshold.Explanation = s })
	field.SetMinimumTextWidthUsing(prototypeMinNameWidth)
	field.Tooltip = unison.NewTooltipWithText(i18n.Text("A explanation of the effects of the threshold state"))
	content.AddChild(field)

	return content
}

func (p *thresholdSettingsPanel) createOpCheckBox(op model.ThresholdOp) *CheckBox {
	c := NewCheckBox(p.pool.dockable.targetMgr, p.threshold.KeyPrefix+op.Key(), op.String(),
		func() unison.CheckState { return unison.CheckStateFromBool(p.threshold.ContainsOp(op)) },
		func(state unison.CheckState) {
			if state == unison.OnCheckState {
				p.threshold.AddOp(op)
			} else {
				p.threshold.RemoveOp(op)
			}
		})
	c.Tooltip = unison.NewTooltipWithText(op.AltString())
	return c
}
