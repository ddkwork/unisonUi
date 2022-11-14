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
	"github.com/richardwilkes/gcs/v5/model/gurps"
	"github.com/richardwilkes/gcs/v5/model/gurps/gid"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/unison"
)

var (
	conditionalModifierColMap = map[int]int{
		0: gurps.ConditionalModifierValueColumn,
		1: gurps.ConditionalModifierDescriptionColumn,
	}
	_ TableProvider[*gurps.ConditionalModifier] = &condModProvider{}
)

type condModProvider struct {
	table    *unison.Table[*Node[*gurps.ConditionalModifier]]
	provider gurps.ConditionalModifierListProvider
}

// NewConditionalModifiersProvider creates a new table provider for conditional modifiers.
func NewConditionalModifiersProvider(provider gurps.ConditionalModifierListProvider) TableProvider[*gurps.ConditionalModifier] {
	return &condModProvider{
		provider: provider,
	}
}

func (p *condModProvider) RefKey() string {
	return gurps.BlockLayoutConditionalModifiersKey
}

func (p *condModProvider) AllTags() []string {
	return nil
}

func (p *condModProvider) SetTable(table *unison.Table[*Node[*gurps.ConditionalModifier]]) {
	p.table = table
}

func (p *condModProvider) RootRowCount() int {
	return len(p.provider.ConditionalModifiers())
}

func (p *condModProvider) RootRows() []*Node[*gurps.ConditionalModifier] {
	data := p.provider.ConditionalModifiers()
	rows := make([]*Node[*gurps.ConditionalModifier], 0, len(data))
	for _, one := range data {
		rows = append(rows, NewNode[*gurps.ConditionalModifier](p.table, nil, conditionalModifierColMap, one, true))
	}
	return rows
}

func (p *condModProvider) SetRootRows(_ []*Node[*gurps.ConditionalModifier]) {
}

func (p *condModProvider) RootData() []*gurps.ConditionalModifier {
	return p.provider.ConditionalModifiers()
}

func (p *condModProvider) SetRootData(_ []*gurps.ConditionalModifier) {
}

func (p *condModProvider) Entity() *gurps.Entity {
	return p.provider.Entity()
}

func (p *condModProvider) DragKey() string {
	return gid.ConditionalModifier
}

func (p *condModProvider) DragSVG() *unison.SVG {
	return nil
}

func (p *condModProvider) DropShouldMoveData(_, _ *unison.Table[*Node[*gurps.ConditionalModifier]]) bool {
	// Not used
	return false
}

func (p *condModProvider) ProcessDropData(_, _ *unison.Table[*Node[*gurps.ConditionalModifier]]) {
}

func (p *condModProvider) AltDropSupport() *AltDropSupport {
	return nil
}

func (p *condModProvider) ItemNames() (singular, plural string) {
	return i18n.Text("Conditional Modifier"), i18n.Text("Conditional Modifiers")
}

func (p *condModProvider) Headers() []unison.TableColumnHeader[*Node[*gurps.ConditionalModifier]] {
	var headers []unison.TableColumnHeader[*Node[*gurps.ConditionalModifier]]
	for i := 0; i < len(conditionalModifierColMap); i++ {
		switch conditionalModifierColMap[i] {
		case gurps.ConditionalModifierValueColumn:
			headers = append(headers, NewEditorListHeader[*gurps.ConditionalModifier]("±", i18n.Text("Modifier"), true))
		case gurps.ConditionalModifierDescriptionColumn:
			headers = append(headers, NewEditorListHeader[*gurps.ConditionalModifier](i18n.Text("Condition"), "", true))
		default:
			jot.Fatalf(1, "invalid conditional modifier column: %d", conditionalModifierColMap[i])
		}
	}
	return headers
}

func (p *condModProvider) SyncHeader(_ []unison.TableColumnHeader[*Node[*gurps.ConditionalModifier]]) {
}

func (p *condModProvider) HierarchyColumnIndex() int {
	return -1
}

func (p *condModProvider) ExcessWidthColumnIndex() int {
	for k, v := range conditionalModifierColMap {
		if v == gurps.ConditionalModifierDescriptionColumn {
			return k
		}
	}
	return 0
}

func (p *condModProvider) OpenEditor(_ Rebuildable, _ *unison.Table[*Node[*gurps.ConditionalModifier]]) {
}

func (p *condModProvider) CreateItem(_ Rebuildable, _ *unison.Table[*Node[*gurps.ConditionalModifier]], _ ItemVariant) {
}

func (p *condModProvider) Serialize() ([]byte, error) {
	return nil, errs.New("not allowed")
}

func (p *condModProvider) Deserialize(_ []byte) error {
	return errs.New("not allowed")
}

func (p *condModProvider) ContextMenuItems() []ContextMenuItem {
	return nil
}
