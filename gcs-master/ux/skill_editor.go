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
	"strings"

	"github.com/richardwilkes/gcs/v5/model/fxp"
	"github.com/richardwilkes/gcs/v5/model/gurps"
	"github.com/richardwilkes/gcs/v5/model/gurps/gid"
	"github.com/richardwilkes/gcs/v5/model/gurps/skill"
	"github.com/richardwilkes/gcs/v5/model/gurps/weapon"
	"github.com/richardwilkes/gcs/v5/svg"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/unison"
)

// EditSkill displays the editor for an skill.
func EditSkill(owner Rebuildable, skill *gurps.Skill) {
	displayEditor[*gurps.Skill, *gurps.SkillEditData](owner, skill, svg.GCSSkills, initSkillEditor)
}

func initSkillEditor(e *editor[*gurps.Skill, *gurps.SkillEditData], content *unison.Panel) func() {
	owner := e.owner.AsPanel().Self
	_, ownerIsSheet := owner.(*Sheet)
	_, ownerIsTemplate := owner.(*Template)
	addNameLabelAndField(content, &e.editorData.Name)
	isTechnique := strings.HasPrefix(e.target.Type, gid.Technique)
	if !e.target.Container() && !isTechnique {
		addSpecializationLabelAndField(content, &e.editorData.Specialization)
		addTechLevelRequired(content, &e.editorData.TechLevel, ownerIsSheet)
	}
	addNotesLabelAndField(content, &e.editorData.LocalNotes)
	addVTTNotesLabelAndField(content, &e.editorData.VTTNotes)
	addTagsLabelAndField(content, &e.editorData.Tags)
	if e.target.Container() {
		addTemplateChoices(content, nil, "", &e.editorData.TemplatePicker)
	} else {
		if isTechnique {
			wrapper := addFlowWrapper(content, i18n.Text("Defaults To"), 4)
			wrapper.SetLayoutData(&unison.FlexLayoutData{
				HAlign: unison.FillAlignment,
				HGrab:  true,
			})
			flags := gurps.TenFlag
			if isTechnique {
				flags |= gurps.SkillFlag + gurps.ParryFlag + gurps.BlockFlag + gurps.DodgeFlag
			}
			choices, attrChoice := gurps.AttributeChoices(e.target.Entity, "", flags, e.editorData.TechniqueDefault.DefaultType)
			attrChoicePopup := addPopup(wrapper, choices, &attrChoice)
			skillDefNameField := addStringField(wrapper, i18n.Text("Technique Default Skill Name"),
				i18n.Text("Skill Name"), &e.editorData.TechniqueDefault.Name)
			skillDefNameField.Watermark = i18n.Text("Skill")
			skillDefNameField.SetLayoutData(&unison.FlexLayoutData{
				HAlign: unison.FillAlignment,
				HGrab:  true,
			})
			skillDefSpecialtyField := addStringField(wrapper, i18n.Text("Technique Default Skill Specialization"),
				i18n.Text("Skill Specialization"), &e.editorData.TechniqueDefault.Specialization)
			skillDefSpecialtyField.Watermark = i18n.Text("Specialization")
			skillDefSpecialtyField.SetLayoutData(&unison.FlexLayoutData{
				HAlign: unison.FillAlignment,
				HGrab:  true,
			})
			lastWasSkillBased := skill.DefaultTypeIsSkillBased(e.editorData.TechniqueDefault.DefaultType)
			if !lastWasSkillBased {
				skillDefNameField.RemoveFromParent()
				skillDefSpecialtyField.RemoveFromParent()
			}
			addDecimalField(wrapper, nil, "", i18n.Text("Technique Default Adjustment"),
				i18n.Text("Default Adjustment"), &e.editorData.TechniqueDefault.Modifier, -fxp.NinetyNine,
				fxp.NinetyNine)
			attrChoicePopup.SelectionCallback = func(_ int, item *gurps.AttributeChoice) {
				e.editorData.TechniqueDefault.DefaultType = item.Key
				if skillBased := skill.DefaultTypeIsSkillBased(e.editorData.TechniqueDefault.DefaultType); skillBased != lastWasSkillBased {
					lastWasSkillBased = skillBased
					if skillBased {
						wrapper.AddChildAtIndex(skillDefNameField, len(wrapper.Children())-1)
						wrapper.AddChildAtIndex(skillDefSpecialtyField, len(wrapper.Children())-1)
					} else {
						skillDefNameField.RemoveFromParent()
						skillDefSpecialtyField.RemoveFromParent()
					}
				}
				MarkModified(content)
			}
			wrapper2 := addFlowWrapper(content, "", 2)
			limitField := NewDecimalField(nil, "", i18n.Text("Limit"),
				func() fxp.Int {
					if e.editorData.TechniqueLimitModifier != nil {
						return *e.editorData.TechniqueLimitModifier
					}
					return 0
				}, func(value fxp.Int) {
					if e.editorData.TechniqueLimitModifier != nil {
						*e.editorData.TechniqueLimitModifier = value
					}
					MarkModified(wrapper2)
				}, -fxp.NinetyNine, fxp.NinetyNine, false, false)
			wrapper2.AddChild(NewCheckBox(nil, "", i18n.Text("Cannot exceed default skill level by more than"),
				func() unison.CheckState {
					return unison.CheckStateFromBool(e.editorData.TechniqueLimitModifier != nil)
				},
				func(state unison.CheckState) {
					if state == unison.OnCheckState {
						if e.editorData.TechniqueLimitModifier == nil {
							var limit fxp.Int
							e.editorData.TechniqueLimitModifier = &limit
						}
						adjustFieldBlank(limitField, false)
					} else {
						e.editorData.TechniqueLimitModifier = nil
						adjustFieldBlank(limitField, true)
					}
				}))
			adjustFieldBlank(limitField, e.editorData.TechniqueLimitModifier == nil)
			wrapper2.AddChild(limitField)
			addLabelAndPopup(content, i18n.Text("Difficulty"), "", skill.AllTechniqueDifficulty,
				&e.editorData.Difficulty.Difficulty)
		} else {
			addDifficultyLabelAndFields(content, e.target.Entity, &e.editorData.Difficulty)
			encLabel := i18n.Text("Encumbrance Penalty")
			wrapper := addFlowWrapper(content, encLabel, 2)
			addDecimalField(wrapper, nil, "", encLabel, "", &e.editorData.EncumbrancePenaltyMultiplier, 0, fxp.Nine)
			wrapper.AddChild(NewFieldTrailingLabel(i18n.Text("times the current encumbrance level")))
		}

		if ownerIsSheet || ownerIsTemplate {
			pointsLabel := i18n.Text("Points")
			wrapper := addFlowWrapper(content, pointsLabel, 3)
			addDecimalField(wrapper, nil, "", pointsLabel, "", &e.editorData.Points, 0, fxp.MaxBasePoints)
			wrapper.AddChild(NewFieldInteriorLeadingLabel(i18n.Text("Level")))
			levelField := NewNonEditableField(func(field *NonEditableField) {
				points := gurps.AdjustedPointsForNonContainerSkillOrTechnique(e.target.Entity, e.editorData.Points,
					e.editorData.Name, e.editorData.Specialization, e.editorData.Tags, nil)
				var level skill.Level
				if isTechnique {
					level = gurps.CalculateTechniqueLevel(e.target.Entity, e.editorData.Name,
						e.editorData.Specialization, e.editorData.Tags, e.editorData.TechniqueDefault,
						e.editorData.Difficulty.Difficulty, points, true, e.editorData.TechniqueLimitModifier)
				} else {
					level = gurps.CalculateSkillLevel(e.target.Entity, e.editorData.Name, e.editorData.Specialization,
						e.editorData.Tags, e.editorData.DefaultedFrom, e.editorData.Difficulty, points,
						e.editorData.EncumbrancePenaltyMultiplier)
				}
				lvl := level.Level.Trunc()
				if lvl <= 0 {
					field.Text = "-"
				} else {
					rsl := level.RelativeLevel
					if isTechnique {
						rsl += e.editorData.TechniqueDefault.Modifier
					}
					field.Text = lvl.String() + "/" + gurps.FormatRelativeSkill(e.target.Entity, e.target.Type,
						e.editorData.Difficulty, rsl)
				}
				field.MarkForLayoutAndRedraw()
			})
			insets := levelField.Border().Insets()
			levelField.SetLayoutData(&unison.FlexLayoutData{
				MinSize: unison.NewSize(levelField.Font.SimpleWidth((-fxp.MaxBasePoints*2).String())+insets.Left+insets.Right, 0),
			})
			wrapper.AddChild(levelField)
		}
	}
	addPageRefLabelAndField(content, &e.editorData.PageRef)
	if !e.target.Container() {
		content.AddChild(newPrereqPanel(e.target.Entity, &e.editorData.Prereq))
		content.AddChild(newDefaultsPanel(e.target.Entity, &e.editorData.Defaults))
		content.AddChild(newFeaturesPanel(e.target.Entity, e.target, &e.editorData.Features))
		for _, wt := range weapon.AllType {
			content.AddChild(newWeaponsPanel(e, e.target, wt, &e.editorData.Weapons))
		}
		content.AddChild(newStudyPanel(e.target.Entity, &e.editorData.Study))
	}
	return nil
}
