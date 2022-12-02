package ux

import (
	"github.com/richardwilkes/unison"
)

// NewFieldLeadingLabel creates a new label appropriate for the first label in a row before a field.
func NewFieldLeadingLabel(text string) *unison.Label {
	label := unison.NewLabel()
	label.Text = text
	label.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.EndAlignment,
		VAlign: unison.MiddleAlignment,
	})
	return label
}

// NewFieldInteriorLeadingLabel creates a new label appropriate for the label in the interior of a row before a field.
func NewFieldInteriorLeadingLabel(text string) *unison.Label {
	label := unison.NewLabel()
	label.Text = text
	label.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.EndAlignment,
		VAlign: unison.MiddleAlignment,
	})
	label.SetBorder(unison.NewEmptyBorder(unison.Insets{Left: unison.StdHSpacing}))
	return label
}

// NewFieldTrailingLabel creates a new label appropriate for after a field.
func NewFieldTrailingLabel(text string) *unison.Label {
	label := unison.NewLabel()
	label.Text = text
	label.SetLayoutData(&unison.FlexLayoutData{
		VAlign: unison.MiddleAlignment,
	})
	return label
}
