package ux

import (
	"github.com/richardwilkes/gcs/v5/model"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/unison"
)

// TableUndoEditData holds the data necessary to provide undo for a table.
type TableUndoEditData[T model.NodeTypes] struct {
	Table *unison.Table[*Node[T]]
	Data  PreservedTableData[T]
}

// NewTableUndoEditData collects the undo edit data for a table.
func NewTableUndoEditData[T model.NodeTypes](table *unison.Table[*Node[T]]) *TableUndoEditData[T] {
	if table == nil {
		return nil
	}
	undo := &TableUndoEditData[T]{Table: table}
	if err := undo.Data.Collect(table); err != nil {
		jot.Error(err)
		return nil
	}
	return undo
}

// Apply the undo edit data to a table.
func (t *TableUndoEditData[T]) Apply() {
	if t == nil {
		return
	}
	if err := t.Data.Apply(t.Table); err != nil {
		jot.Error(err)
	}
}

// TableDragUndoEditData holds the undo edit data for a table drag.
type TableDragUndoEditData[T model.NodeTypes] struct {
	From *TableUndoEditData[T]
	To   *TableUndoEditData[T]
}

// NewTableDragUndoEditData collects the undo edit data for a table drag.
func NewTableDragUndoEditData[T model.NodeTypes](from, to *unison.Table[*Node[T]]) *TableDragUndoEditData[T] {
	return &TableDragUndoEditData[T]{
		From: NewTableUndoEditData(from),
		To:   NewTableUndoEditData(to),
	}
}

// Apply the undo edit data to a table.
func (t *TableDragUndoEditData[T]) Apply() {
	t.To.Apply()
	if t.From != nil {
		t.From.Apply()
	}
}
