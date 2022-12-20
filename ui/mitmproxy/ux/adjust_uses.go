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
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/unison"
)

type adjustUsesListUndoEdit = *unison.UndoEdit[*adjustUsesList]

type adjustUsesList struct {
	Owner Rebuildable
	List  []*usesAdjuster
}

func (a *adjustUsesList) Apply() {
	for _, one := range a.List {
		one.Apply()
	}
	MarkModified(a.Owner)
}

type usesAdjuster struct {
	Target *model.Equipment
	Uses   int
}

func newUsesAdjuster(target *model.Equipment) *usesAdjuster {
	return &usesAdjuster{
		Target: target,
		Uses:   target.Uses,
	}
}

func (a *usesAdjuster) Apply() {
	a.Target.Uses = a.Uses
}

func canAdjustUses(table *unison.Table[*Node[*model.Equipment]], amount int) bool {
	for _, row := range table.SelectedRows(false) {
		if eqp := row.Data(); eqp != nil {
			total := eqp.Uses + amount
			if total >= 0 && total <= eqp.MaxUses {
				return true
			}
		}
	}
	return false
}

func adjustUses(owner Rebuildable, table *unison.Table[*Node[*model.Equipment]], amount int) {
	before := &adjustUsesList{Owner: owner}
	after := &adjustUsesList{Owner: owner}
	for _, row := range table.SelectedRows(false) {
		if eqp := row.Data(); eqp != nil {
			total := eqp.Uses + amount
			if total >= 0 && total <= eqp.MaxUses {
				before.List = append(before.List, newUsesAdjuster(eqp))
				eqp.Uses = total
				after.List = append(after.List, newUsesAdjuster(eqp))
			}
		}
	}
	if len(before.List) > 0 {
		if mgr := unison.UndoManagerFor(table); mgr != nil {
			var name string
			if amount < 0 {
				name = i18n.Text("Decrease Uses")
			} else {
				name = i18n.Text("Increase Uses")
			}
			mgr.Add(&unison.UndoEdit[*adjustUsesList]{
				ID:         unison.NextUndoID(),
				EditName:   name,
				UndoFunc:   func(edit adjustUsesListUndoEdit) { edit.BeforeData.Apply() },
				RedoFunc:   func(edit adjustUsesListUndoEdit) { edit.AfterData.Apply() },
				BeforeData: before,
				AfterData:  after,
			})
		}
		MarkModified(before.Owner)
	}
}
