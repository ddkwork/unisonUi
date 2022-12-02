package model

import (
	"github.com/richardwilkes/toolbox/i18n"
)

// TraitModifierData holds the TraitModifier data that is written to disk.
type TraitModifierData struct {
	ContainerBase[*TraitModifier]
	TraitModifierEditData
}

// Kind returns the kind of data.
func (d *TraitModifierData) Kind() string {
	return d.kind(i18n.Text("Trait Modifier"))
}

// ClearUnusedFieldsForType zeroes out the fields that are not applicable to this type (container vs not-container).
func (d *TraitModifierData) ClearUnusedFieldsForType() {
	d.clearUnusedFields()
	if d.Container() {
		d.CostType = 0
		d.Disabled = false
		d.Cost = 0
		d.Levels = 0
		d.Affects = 0
		d.Features = nil
	}
}
