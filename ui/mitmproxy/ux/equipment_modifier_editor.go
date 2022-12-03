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

const equipmentCostAndWeightPrototype = "-99.99 CF"

// EditEquipmentModifier displays the editor for an equipment modifier.
func EditEquipmentModifier(owner Rebuildable, modifier *model.EquipmentModifier) {
	displayEditor[*model.EquipmentModifier, *model.EquipmentModifierEditData](owner, modifier,
		svg.GCSEquipmentModifiers, nil, initEquipmentModifierEditor)
}

func initEquipmentModifierEditor(e *editor[*model.EquipmentModifier, *model.EquipmentModifierEditData], content *unison.Panel) func() {
	addNameLabelAndField(content, &e.editorData.Name)
	if !e.target.Container() {
		addLabelAndStringField(content, i18n.Text("Tech Level"), model.TechLevelInfo, &e.editorData.TechLevel)
	}
	addNotesLabelAndField(content, &e.editorData.LocalNotes)
	addVTTNotesLabelAndField(content, &e.editorData.VTTNotes)
	if !e.target.Container() {
		content.AddChild(unison.NewPanel())
		addInvertedCheckBox(content, i18n.Text("Enabled"), &e.editorData.Disabled)
		addEquipmentCostFields(content, e)
		addEquipmentWeightFields(content, e)
	}
	addTagsLabelAndField(content, &e.editorData.Tags)
	addPageRefLabelAndField(content, &e.editorData.PageRef)
	if !e.target.Container() {
		content.AddChild(newFeaturesPanel(e.target.Entity, e.target, &e.editorData.Features))
	}
	return nil
}

func addEquipmentCostFields(parent *unison.Panel, e *editor[*model.EquipmentModifier, *model.EquipmentModifierEditData]) {
	label := i18n.Text("Cost Modifier")
	wrapper := addFlowWrapper(parent, label, 2)
	field := NewStringField(nil, "", label,
		func() string { return e.editorData.CostType.Format(e.editorData.CostAmount) },
		func(value string) {
			e.editorData.CostAmount = e.editorData.CostType.Format(value)
			MarkModified(parent)
		})
	field.SetMinimumTextWidthUsing(equipmentCostAndWeightPrototype)
	wrapper.AddChild(field)
	popup := unison.NewPopupMenu[string]()
	for _, one := range model.AllEquipmentModifierCostType {
		popup.AddItem(one.StringWithExample())
	}
	popup.SelectIndex(int(e.editorData.CostType))
	wrapper.AddChild(popup)
	popup.SelectionCallback = func(index int, _ string) {
		e.editorData.CostType = model.AllEquipmentModifierCostType[index]
		field.SetText(e.editorData.CostType.Format(field.Text()))
		MarkModified(wrapper)
	}
}

func addEquipmentWeightFields(parent *unison.Panel, e *editor[*model.EquipmentModifier, *model.EquipmentModifierEditData]) {
	units := model.SheetSettingsFor(e.target.Entity).DefaultWeightUnits
	label := i18n.Text("Weight Modifier")
	wrapper := addFlowWrapper(parent, label, 2)
	field := NewStringField(nil, "", label,
		func() string { return e.editorData.WeightType.Format(e.editorData.WeightAmount, units) },
		func(value string) {
			e.editorData.WeightAmount = e.editorData.WeightType.Format(value, units)
			MarkModified(parent)
		})
	field.SetMinimumTextWidthUsing(equipmentCostAndWeightPrototype)
	wrapper.AddChild(field)
	popup := unison.NewPopupMenu[string]()
	for _, one := range model.AllEquipmentModifierWeightType {
		popup.AddItem(one.StringWithExample())
	}
	popup.SelectIndex(int(e.editorData.WeightType))
	wrapper.AddChild(popup)
	popup.SelectionCallback = func(index int, _ string) {
		e.editorData.WeightType = model.AllEquipmentModifierWeightType[index]
		field.SetText(e.editorData.WeightType.Format(field.Text(), units))
		MarkModified(wrapper)
	}
}
