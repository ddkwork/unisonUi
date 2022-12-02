package ux

import (
	"github.com/richardwilkes/unison"
)

// StringField holds the value for a string field.
type StringField struct {
	*unison.Field
	undoID    int64
	targetMgr *TargetMgr
	targetKey string
	undoTitle string
	last      string
	get       func() string
	set       func(string)
	useGet    bool
	inUndo    bool
}

// NewMultiLineStringField creates a new field for editing a string.
func NewMultiLineStringField(targetMgr *TargetMgr, targetKey, undoTitle string, get func() string, set func(string)) *StringField {
	f := unison.NewMultiLineField()
	f.SetWrap(true)
	return newStringField(f, targetMgr, targetKey, undoTitle, get, set)
}

// NewStringField creates a new field for editing a string.
func NewStringField(targetMgr *TargetMgr, targetKey, undoTitle string, get func() string, set func(string)) *StringField {
	return newStringField(unison.NewField(), targetMgr, targetKey, undoTitle, get, set)
}

func newStringField(field *unison.Field, targetMgr *TargetMgr, targetKey, undoTitle string, get func() string, set func(string)) *StringField {
	f := &StringField{
		Field:     field,
		undoID:    unison.NextUndoID(),
		targetMgr: targetMgr,
		targetKey: targetKey,
		undoTitle: undoTitle,
		last:      get(),
		get:       get,
		set:       set,
		useGet:    true,
	}
	f.Self = f
	f.LostFocusCallback = f.lostFocus
	f.ModifiedCallback = f.modified
	f.Sync()
	f.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		HGrab:  true,
	})
	if targetMgr != nil && targetKey != "" {
		f.RefKey = targetKey
	}
	return f
}

func (f *StringField) lostFocus() {
	f.useGet = true
	f.SetText(f.Text())
	f.DefaultFocusLost()
}

func (f *StringField) getData() string {
	if f.useGet {
		f.useGet = false
		return f.get()
	}
	return f.Text()
}

func (f *StringField) modified() {
	text := f.Text()
	if !f.inUndo && f.undoID != unison.NoUndoID {
		if mgr := unison.UndoManagerFor(f); mgr != nil {
			undo := NewTargetUndo(f.targetMgr, f.targetKey, f.undoTitle, f.undoID, func(target *unison.Panel, data string) {
				self := f
				if target != nil {
					if field, ok := target.Self.(*StringField); ok {
						self = field
					}
				}
				self.setWithoutUndo(data, true)
			}, f.get())
			undo.AfterData = text
			mgr.Add(undo)
		}
	}
	if f.last != text {
		f.last = text
		f.set(text)
		MarkForLayoutWithinDockable(f)
		MarkModified(f)
	}
}

func (f *StringField) setWithoutUndo(text string, focus bool) {
	f.inUndo = true
	f.SetText(text)
	f.inUndo = false
	if focus {
		f.RequestFocus()
		f.SelectAll()
	}
}

// Sync the field to the current value.
func (f *StringField) Sync() {
	if !f.Focused() {
		f.useGet = true
	}
	f.setWithoutUndo(f.getData(), false)
}
