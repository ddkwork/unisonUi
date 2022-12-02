package model

import (
	"strings"

	"github.com/richardwilkes/toolbox/i18n"
)

// Possible values.
const (
	PercentageTraitModifierCostType TraitModifierCostType = iota
	PointsTraitModifierCostType
	MultiplierTraitModifierCostType
	LastTraitModifierCostType = MultiplierTraitModifierCostType
)

var (
	// AllTraitModifierCostType holds all possible values.
	AllTraitModifierCostType = []TraitModifierCostType{
		PercentageTraitModifierCostType,
		PointsTraitModifierCostType,
		MultiplierTraitModifierCostType,
	}
	traitModifierCostTypeData = []struct {
		key    string
		string string
	}{
		{
			key:    "percentage",
			string: i18n.Text("%"),
		},
		{
			key:    "points",
			string: i18n.Text("points"),
		},
		{
			key:    "multiplier",
			string: i18n.Text("×"),
		},
	}
)

// TraitModifierCostType describes how a TraitModifier's point cost is applied.
type TraitModifierCostType byte

// EnsureValid ensures this is of a known value.
func (enum TraitModifierCostType) EnsureValid() TraitModifierCostType {
	if enum <= LastTraitModifierCostType {
		return enum
	}
	return 0
}

// Key returns the key used in serialization.
func (enum TraitModifierCostType) Key() string {
	return traitModifierCostTypeData[enum.EnsureValid()].key
}

// String implements fmt.Stringer.
func (enum TraitModifierCostType) String() string {
	return traitModifierCostTypeData[enum.EnsureValid()].string
}

// ExtractTraitModifierCostType extracts the value from a string.
func ExtractTraitModifierCostType(str string) TraitModifierCostType {
	for i, one := range traitModifierCostTypeData {
		if strings.EqualFold(one.key, str) {
			return TraitModifierCostType(i)
		}
	}
	return 0
}

// MarshalText implements the encoding.TextMarshaler interface.
func (enum TraitModifierCostType) MarshalText() (text []byte, err error) {
	return []byte(enum.Key()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (enum *TraitModifierCostType) UnmarshalText(text []byte) error {
	*enum = ExtractTraitModifierCostType(string(text))
	return nil
}
