package model

import (
	"github.com/richardwilkes/toolbox/i18n"
)

// TraitData holds the Trait data that is written to disk.
type TraitData struct {
	ContainerBase[*Trait]
	TraitEditData
}

// Kind returns the kind of data.
func (d *TraitData) Kind() string {
	return d.kind(i18n.Text("Trait"))
}

// ClearUnusedFieldsForType zeroes out the fields that are not applicable to this type (container vs not-container).
func (d *TraitData) ClearUnusedFieldsForType() {
	d.clearUnusedFields()
	if d.Container() {
		d.BasePoints = 0
		d.Levels = 0
		d.PointsPerLevel = 0
		d.CanLevel = false
		d.Prereq = nil
		d.Weapons = nil
		d.Features = nil
		d.RoundCostDown = false
		if d.TemplatePicker == nil {
			d.TemplatePicker = &TemplatePicker{}
		}
	} else {
		d.ContainerType = 0
		d.TemplatePicker = nil
		d.Ancestry = ""
		if !d.CanLevel {
			d.Levels = 0
			d.PointsPerLevel = 0
		}
	}
}
