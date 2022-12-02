package model

import (
	"github.com/richardwilkes/gcs/v5/model/fxp"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/toolbox/xio"
)

var _ Prereq = &TraitPrereq{}

// TraitPrereq holds a prereq against a Trait.
type TraitPrereq struct {
	Parent        *PrereqList     `json:"-"`
	Type          PrereqType      `json:"type"`
	Has           bool            `json:"has"`
	NameCriteria  StringCriteria  `json:"name,omitempty"`
	LevelCriteria NumericCriteria `json:"level,omitempty"`
	NotesCriteria StringCriteria  `json:"notes,omitempty"`
}

// NewTraitPrereq creates a new TraitPrereq.
func NewTraitPrereq() *TraitPrereq {
	return &TraitPrereq{
		Type: TraitPrereqType,
		NameCriteria: StringCriteria{
			StringCriteriaData: StringCriteriaData{
				Compare: IsString,
			},
		},
		LevelCriteria: NumericCriteria{
			NumericCriteriaData: NumericCriteriaData{
				Compare: AtLeastNumber,
			},
		},
		NotesCriteria: StringCriteria{
			StringCriteriaData: StringCriteriaData{
				Compare: AnyString,
			},
		},
		Has: true,
	}
}

// PrereqType implements Prereq.
func (a *TraitPrereq) PrereqType() PrereqType {
	return a.Type
}

// ParentList implements Prereq.
func (a *TraitPrereq) ParentList() *PrereqList {
	return a.Parent
}

// Clone implements Prereq.
func (a *TraitPrereq) Clone(parent *PrereqList) Prereq {
	clone := *a
	clone.Parent = parent
	return &clone
}

// FillWithNameableKeys implements Prereq.
func (a *TraitPrereq) FillWithNameableKeys(m map[string]string) {
	Extract(a.NameCriteria.Qualifier, m)
	Extract(a.NotesCriteria.Qualifier, m)
}

// ApplyNameableKeys implements Prereq.
func (a *TraitPrereq) ApplyNameableKeys(m map[string]string) {
	a.NameCriteria.Qualifier = Apply(a.NameCriteria.Qualifier, m)
	a.NotesCriteria.Qualifier = Apply(a.NotesCriteria.Qualifier, m)
}

// Satisfied implements Prereq.
func (a *TraitPrereq) Satisfied(entity *Entity, exclude any, tooltip *xio.ByteBuffer, prefix string, _ *bool) bool {
	satisfied := false
	Traverse(func(t *Trait) bool {
		if exclude == t || !a.NameCriteria.Matches(t.Name) {
			return false
		}
		notes := t.Notes()
		if modNotes := t.ModifierNotes(); modNotes != "" {
			notes += "\n" + modNotes
		}
		if !a.NotesCriteria.Matches(notes) {
			return false
		}
		var levels fxp.Int
		if t.IsLeveled() {
			levels = t.Levels.Max(0)
		}
		satisfied = a.LevelCriteria.Matches(levels)
		return satisfied
	}, true, false, entity.Traits...)
	if !a.Has {
		satisfied = !satisfied
	}
	if !satisfied && tooltip != nil {
		tooltip.WriteString(prefix)
		tooltip.WriteString(HasText(a.Has))
		tooltip.WriteString(i18n.Text(" a trait whose name "))
		tooltip.WriteString(a.NameCriteria.String())
		if a.NotesCriteria.Compare != AnyString {
			tooltip.WriteString(i18n.Text(", notes "))
			tooltip.WriteString(a.NotesCriteria.String())
			tooltip.WriteByte(',')
		}
		tooltip.WriteString(i18n.Text(" and level "))
		tooltip.WriteString(a.LevelCriteria.String())
	}
	return satisfied
}
