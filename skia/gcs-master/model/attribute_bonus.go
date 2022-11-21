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

package model

import (
	"fmt"

	"github.com/richardwilkes/gcs/v5/model/fxp"
	"github.com/richardwilkes/toolbox/xio"
)

var _ Bonus = &AttributeBonus{}

// AttributeBonus holds the data for a bonus to an attribute.
type AttributeBonus struct {
	Type       FeatureType     `json:"type"`
	Limitation BonusLimitation `json:"limitation,omitempty"`
	Attribute  string          `json:"attribute"`
	LeveledAmount
	owner fmt.Stringer
}

// NewAttributeBonus creates a new AttributeBonus.
func NewAttributeBonus(attrID string) *AttributeBonus {
	return &AttributeBonus{
		Type:          AttributeBonusFeatureType,
		Attribute:     attrID,
		Limitation:    NoneBonusLimitation,
		LeveledAmount: LeveledAmount{Amount: fxp.One},
	}
}

// FeatureType implements Feature.
func (a *AttributeBonus) FeatureType() FeatureType {
	return a.Type
}

// Clone implements Feature.
func (a *AttributeBonus) Clone() Feature {
	other := *a
	return &other
}

// FillWithNameableKeys implements Feature.
func (a *AttributeBonus) FillWithNameableKeys(_ map[string]string) {
}

// ApplyNameableKeys implements Feature.
func (a *AttributeBonus) ApplyNameableKeys(_ map[string]string) {
}

// Owner implements Bonus.
func (a *AttributeBonus) Owner() fmt.Stringer {
	return a.owner
}

// SetOwner implements Bonus.
func (a *AttributeBonus) SetOwner(owner fmt.Stringer) {
	a.owner = owner
}

// SetLevel implements Bonus.
func (a *AttributeBonus) SetLevel(level fxp.Int) {
	a.Level = level
}

// AddToTooltip implements Bonus.
func (a *AttributeBonus) AddToTooltip(buffer *xio.ByteBuffer) {
	basicAddToTooltip(a.owner, &a.LeveledAmount, buffer)
}
