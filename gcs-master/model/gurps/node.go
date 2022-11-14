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

package gurps

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/richardwilkes/gcs/v5/model/fxp"
)

// NodeTypes is a constraint that defines the types that may be nodes.
type NodeTypes interface {
	*ConditionalModifier | *Equipment | *EquipmentModifier | *Note | *Skill | *Spell | *Trait | *TraitModifier | *Weapon
}

// Node defines the methods required of nodes in our tables.
type Node[T NodeTypes] interface {
	fmt.Stringer
	UUID() uuid.UUID
	Clone(newEntity *Entity, newParent T, preserveID bool) T
	OwningEntity() *Entity
	SetOwningEntity(entity *Entity)
	Kind() string
	Container() bool
	Parent() T
	SetParent(parent T)
	HasChildren() bool
	NodeChildren() []T
	SetChildren(children []T)
	Enabled() bool
	Open() bool
	SetOpen(open bool)
	CellData(column int, data *CellData)
	FillWithNameableKeys(m map[string]string)
	ApplyNameableKeys(m map[string]string)
}

// RawPointsAdjuster defines methods for nodes that can have their raw points adjusted must implement.
type RawPointsAdjuster[T NodeTypes] interface {
	Node[T]
	RawPoints() fxp.Int
	SetRawPoints(points fxp.Int) bool
}

// SkillAdjustmentProvider defines methods for nodes that can have their skill level adjusted must implement.
type SkillAdjustmentProvider[T NodeTypes] interface {
	RawPointsAdjuster[T]
	IncrementSkillLevel()
	DecrementSkillLevel()
}

// EditorData defines the methods required of editor data.
type EditorData[T NodeTypes] interface {
	// CopyFrom copies the corresponding data from the node into this editor data.
	CopyFrom(T)
	// ApplyTo copes he editor data into the provided node.
	ApplyTo(T)
}

// AsNode converts a T to a Node[T]. This shouldn't require these hoops, but Go generics (as of 1.19) fails to compile
// otherwise.
func AsNode[T NodeTypes](in T) Node[T] {
	return any(in).(Node[T])
}

// CloneNodes creates clones of the provided nodes.
func CloneNodes[T NodeTypes](newEntity *Entity, newParent T, preserveID bool, nodes []Node[T]) []T {
	clones := make([]T, len(nodes))
	for i, one := range nodes {
		clones[i] = one.Clone(newEntity, newParent, preserveID)
	}
	return clones
}
