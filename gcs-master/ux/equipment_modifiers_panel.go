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
	"github.com/google/uuid"
	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/unison"
)

type equipmentModifiersPanel struct {
	unison.Panel
	entity    *model.Entity
	modifiers *[]*model.EquipmentModifier
	provider  TableProvider[*model.EquipmentModifier]
	table     *unison.Table[*Node[*model.EquipmentModifier]]
}

func newEquipmentModifiersPanel(entity *model.Entity, modifiers *[]*model.EquipmentModifier) *equipmentModifiersPanel {
	p := &equipmentModifiersPanel{
		entity:    entity,
		modifiers: modifiers,
	}
	p.Self = p
	p.SetLayout(&unison.FlexLayout{Columns: 1})
	p.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  2,
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})
	p.SetBorder(unison.NewLineBorder(model.HeaderColor, 0, unison.NewUniformInsets(1), false))
	p.DrawCallback = func(gc *unison.Canvas, rect unison.Rect) {
		gc.DrawRect(rect, unison.ContentColor.Paint(gc, rect, unison.Fill))
	}
	p.provider = NewEquipmentModifiersProvider(p, true)
	p.table = newEditorTable(p.AsPanel(), p.provider)
	p.table.RefKey = "equipment-modifiers-" + uuid.New().String()
	return p
}

func (p *equipmentModifiersPanel) Entity() *model.Entity {
	return p.entity
}

func (p *equipmentModifiersPanel) EquipmentModifierList() []*model.EquipmentModifier {
	return *p.modifiers
}

func (p *equipmentModifiersPanel) SetEquipmentModifierList(list []*model.EquipmentModifier) {
	*p.modifiers = list
	sel := p.table.CopySelectionMap()
	p.table.SyncToModel()
	p.table.SetSelectionMap(sel)
}
