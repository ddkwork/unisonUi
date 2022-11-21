// Copyright ©2021-2022 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package unison

import (
	"time"

	"github.com/richardwilkes/toolbox"
	"github.com/richardwilkes/toolbox/xmath"
)

// DefaultListTheme holds the default ListTheme values for Lists. Modifying this data will not alter existing Lists,
// but will alter any Lists created in the future.
var DefaultListTheme = ListTheme{
	BackgroundInk:          ContentColor,
	OnBackgroundInk:        OnContentColor,
	BandingInk:             BandingColor,
	OnBandingInk:           OnBandingColor,
	SelectionInk:           SelectionColor,
	OnSelectionInk:         OnSelectionColor,
	InactiveSelectionInk:   InactiveSelectionColor,
	OnInactiveSelectionInk: OnInactiveSelectionColor,
	FlashAnimationTime:     100 * time.Millisecond,
}

// ListTheme holds theming data for a List.
type ListTheme struct {
	BackgroundInk          Ink
	OnBackgroundInk        Ink
	BandingInk             Ink
	OnBandingInk           Ink
	SelectionInk           Ink
	OnSelectionInk         Ink
	InactiveSelectionInk   Ink
	OnInactiveSelectionInk Ink
	FlashAnimationTime     time.Duration
}

// List provides a control that allows the user to select from a list of items, represented by cells.
type List[T any] struct {
	Panel
	ListTheme
	DoubleClickCallback  func()
	NewSelectionCallback func()
	Factory              CellFactory
	rows                 []T
	Selection            *xmath.BitSet
	savedSelection       *xmath.BitSet
	anchor               int
	lastSel              int
	allowMultiple        bool
	pressed              bool
	suppressSelection    bool
	suppressScroll       bool
	wasDragged           bool
}

// NewList creates a new List control.
func NewList[T any]() *List[T] {
	l := &List[T]{
		ListTheme:      DefaultListTheme,
		Factory:        &DefaultCellFactory{},
		Selection:      &xmath.BitSet{},
		savedSelection: &xmath.BitSet{},
		anchor:         -1,
		lastSel:        -1,
		allowMultiple:  true,
	}
	l.Self = l
	l.SetFocusable(true)
	l.SetSizer(l.DefaultSizes)
	l.DrawCallback = l.DefaultDraw
	l.GainedFocusCallback = l.DefaultFocusGained
	l.MouseDownCallback = l.DefaultMouseDown
	l.MouseDragCallback = l.DefaultMouseDrag
	l.MouseUpCallback = l.DefaultMouseUp
	l.KeyDownCallback = l.DefaultKeyDown
	l.InstallCmdHandlers(SelectAllItemID, func(_ any) bool { return l.CanSelectAll() }, func(_ any) { l.SelectAll() })
	return l
}

// Count returns the number of rows.
func (l *List[T]) Count() int {
	return len(l.rows)
}

// DataAtIndex returns the data for the specified row index.
func (l *List[T]) DataAtIndex(index int) T {
	if index >= 0 && index < len(l.rows) {
		return l.rows[index]
	}
	var zero T
	return zero
}

// Append values to the list of items.
func (l *List[T]) Append(values ...T) {
	l.rows = append(l.rows, values...)
	l.MarkForLayoutAndRedraw()
}

// Insert values at the specified index.
func (l *List[T]) Insert(index int, values ...T) {
	if index < 0 || index > len(l.rows) {
		index = len(l.rows)
	}
	l.rows = append(l.rows[:index], append(values, l.rows[index:]...)...)
	l.MarkForLayoutAndRedraw()
}

// Replace the value at the specified index.
func (l *List[T]) Replace(index int, value T) {
	if index >= 0 && index < len(l.rows) {
		l.rows[index] = value
		l.MarkForLayoutAndRedraw()
	}
}

// Remove the item at the specified index.
func (l *List[T]) Remove(index int) {
	if index >= 0 && index < len(l.rows) {
		copy(l.rows[index:], l.rows[index+1:])
		size := len(l.rows) - 1
		var zero T
		l.rows[size] = zero
		l.rows = l.rows[:size]
		l.MarkForLayoutAndRedraw()
	}
}

// RemoveRange removes the items at the specified index range, inclusive.
func (l *List[T]) RemoveRange(from, to int) {
	if from >= 0 && from < len(l.rows) && to >= from && to < len(l.rows) {
		copy(l.rows[from:], l.rows[to+1:])
		size := len(l.rows) - (1 + to - from)
		var zero T
		for i := size; i < len(l.rows); i++ {
			l.rows[i] = zero
		}
		l.rows = l.rows[:size]
		l.MarkForLayoutAndRedraw()
	}
}

// DefaultSizes provides the default sizing.
func (l *List[T]) DefaultSizes(hint Size) (min, pref, max Size) {
	max = MaxSize(max)
	height := xmath.Ceil(l.Factory.CellHeight())
	if height < 1 {
		height = 0
	}
	size := Size{Width: hint.Width, Height: height}
	for row := range l.rows {
		cell := l.cell(row)
		_, cPref, cMax := cell.Sizes(size)
		cPref.GrowToInteger()
		cMax.GrowToInteger()
		if pref.Width < cPref.Width {
			pref.Width = cPref.Width
		}
		if max.Width < cMax.Width {
			max.Width = cMax.Width
		}
		if height < 1 {
			pref.Height += cPref.Height
			max.Height += cMax.Height
		}
	}
	if height >= 1 {
		count := float32(len(l.rows))
		if count < 1 {
			count = 1
		}
		pref.Height = count * height
		max.Height = count * height
		if max.Height < DefaultMaxSize {
			max.Height = DefaultMaxSize
		}
	}
	if border := l.Border(); border != nil {
		insets := border.Insets()
		pref.AddInsets(insets)
		max.AddInsets(insets)
	}
	pref.GrowToInteger()
	max.GrowToInteger()
	return pref, pref, max
}

// DefaultFocusGained provides the default focus gained handling.
func (l *List[T]) DefaultFocusGained() {
	if !l.suppressScroll {
		l.ScrollIntoView()
	}
	l.MarkForRedraw()
}

func (l *List[T]) cellParams(row int) (fg, bg Ink, selected, focused bool) {
	focused = l.Focused()
	if !l.suppressSelection {
		selected = l.Selection.State(row)
	}
	switch {
	case selected && focused:
		fg = l.OnSelectionInk
		bg = l.SelectionInk
	case selected:
		fg = l.OnInactiveSelectionInk
		bg = l.InactiveSelectionInk
	case row%2 == 1:
		fg = l.OnBandingInk
		bg = l.BandingInk
	default:
		fg = l.OnBackgroundInk
		bg = l.BackgroundInk
	}
	return fg, bg, selected, focused
}

func (l *List[T]) cell(row int) *Panel {
	fg, bg, selected, focused := l.cellParams(row)
	return l.Factory.CreateCell(l, l.rows[row], row, fg, bg, selected, focused).AsPanel()
}

// DefaultDraw provides the default drawing.
func (l *List[T]) DefaultDraw(canvas *Canvas, dirty Rect) {
	row, y := l.rowAt(dirty.Y)
	if row >= 0 {
		cellHeight := xmath.Ceil(l.Factory.CellHeight())
		count := len(l.rows)
		yMax := dirty.Y + dirty.Height
		rect := l.ContentRect(false)
		for row < count && y < yMax {
			fg, bg, selected, focused := l.cellParams(row)
			cell := l.Factory.CreateCell(l, l.rows[row], row, fg, bg, selected, focused).AsPanel()
			cellRect := Rect{Point: Point{X: rect.X, Y: y}, Size: Size{Width: rect.Width, Height: cellHeight}}
			if cellHeight < 1 {
				_, pref, _ := cell.Sizes(Size{})
				pref.GrowToInteger()
				cellRect.Height = pref.Height
			}
			cell.SetFrameRect(cellRect)
			y += cellRect.Height
			r := NewRect(rect.X, cellRect.Y, rect.Width, cellRect.Height)
			canvas.DrawRect(r, bg.Paint(canvas, r, Fill))
			canvas.Save()
			tl := cellRect.Point
			dirty.Point.Subtract(tl)
			canvas.Translate(cellRect.X, cellRect.Y)
			cellRect.X = 0
			cellRect.Y = 0
			cell.Draw(canvas, dirty)
			dirty.Point.Add(tl)
			canvas.Restore()
			row++
		}
	}
}

// DefaultMouseDown provides the default mouse down handling.
func (l *List[T]) DefaultMouseDown(where Point, button, clickCount int, mod Modifiers) bool {
	l.suppressScroll = true
	l.RequestFocus()
	l.suppressScroll = false
	l.savedSelection = l.Selection.Clone()
	l.lastSel = -1
	l.wasDragged = false
	if index, _ := l.rowAt(where.Y); index >= 0 {
		switch {
		case mod.DiscontiguousSelectionDown():
			if l.allowMultiple {
				l.Selection.Flip(index)
			} else {
				wasSet := l.Selection.State(index)
				l.Selection.Reset()
				if !wasSet {
					l.Selection.Set(index)
				}
			}
			l.anchor = index
		case mod.ShiftDown():
			if l.allowMultiple {
				if l.anchor != -1 {
					l.Selection.SetRange(l.anchor, index)
				} else {
					l.Selection.Set(index)
					l.anchor = index
				}
			} else {
				l.Selection.Reset()
				l.Selection.Set(index)
			}
		case l.Selection.State(index):
			l.lastSel = index
			l.anchor = index
			if clickCount == 2 && l.DoubleClickCallback != nil {
				toolbox.Call(l.DoubleClickCallback)
				return true
			}
		default:
			l.Selection.Reset()
			l.Selection.Set(index)
			l.anchor = index
		}
		if !l.Selection.Equal(l.savedSelection) {
			l.MarkForRedraw()
		}
	}
	l.pressed = true
	return true
}

// DefaultMouseDrag provides the default mouse drag handling.
func (l *List[T]) DefaultMouseDrag(where Point, button int, mod Modifiers) bool {
	if l.pressed {
		l.wasDragged = true
		l.Selection.Copy(l.savedSelection)
		if index, _ := l.rowAt(where.Y); index >= 0 {
			if l.allowMultiple {
				if l.anchor == -1 {
					l.anchor = index
				}
				switch {
				case mod.DiscontiguousSelectionDown():
					l.Selection.FlipRange(l.anchor, index)
				case mod.ShiftDown():
					l.Selection.SetRange(l.anchor, index)
				default:
					l.Selection.Reset()
					l.Selection.SetRange(l.anchor, index)
				}
			} else {
				l.Selection.Reset()
				l.Selection.Set(index)
				l.anchor = index
			}
			if !l.Selection.Equal(l.savedSelection) {
				l.MarkForRedraw()
			}
		}
	}
	return true
}

// DefaultMouseUp provides the default mouse up handling.
func (l *List[T]) DefaultMouseUp(where Point, button int, mod Modifiers) bool {
	if l.pressed {
		l.pressed = false
		if !l.wasDragged && l.lastSel != -1 {
			l.Selection.Reset()
			l.Selection.Set(l.lastSel)
			l.anchor = l.lastSel
			l.MarkForRedraw()
		}
		if l.NewSelectionCallback != nil && !l.Selection.Equal(l.savedSelection) {
			toolbox.Call(l.NewSelectionCallback)
		}
	}
	l.savedSelection = nil
	return true
}

// DefaultKeyDown provides the default key down handling.
func (l *List[T]) DefaultKeyDown(keyCode KeyCode, mod Modifiers, repeat bool) bool {
	if IsControlAction(keyCode, mod) {
		if l.DoubleClickCallback != nil && l.Selection.Count() > 0 {
			toolbox.Call(l.DoubleClickCallback)
		}
		return true
	}
	switch keyCode {
	case KeyUp:
		var first int
		if l.Selection.Count() == 0 {
			first = len(l.rows) - 1
		} else {
			first = l.Selection.FirstSet() - 1
			if first < 0 {
				first = 0
			}
		}
		l.Select(mod.ShiftDown(), first)
		if l.NewSelectionCallback != nil {
			toolbox.Call(l.NewSelectionCallback)
		}
	case KeyDown:
		last := l.Selection.LastSet() + 1
		if last >= len(l.rows) {
			last = len(l.rows) - 1
		}
		l.Select(mod.ShiftDown(), last)
		if l.NewSelectionCallback != nil {
			toolbox.Call(l.NewSelectionCallback)
		}
	case KeyHome:
		l.Select(mod.ShiftDown(), 0)
		if l.NewSelectionCallback != nil {
			toolbox.Call(l.NewSelectionCallback)
		}
	case KeyEnd:
		l.Select(mod.ShiftDown(), len(l.rows)-1)
		if l.NewSelectionCallback != nil {
			toolbox.Call(l.NewSelectionCallback)
		}
	default:
		return false
	}
	return true
}

// CanSelectAll returns true if the list's selection can be expanded.
func (l *List[T]) CanSelectAll() bool {
	return l.Selection.Count() < len(l.rows)
}

// SelectAll selects all of the rows in the list.
func (l *List[T]) SelectAll() {
	l.SelectRange(0, len(l.rows)-1, false)
}

// SelectRange selects items from 'start' to 'end', inclusive. If 'add' is true, then any existing selection is added to
// rather than replaced.
func (l *List[T]) SelectRange(start, end int, add bool) {
	if !l.allowMultiple {
		add = false
		end = start
	}
	if !add {
		l.Selection.Reset()
		l.anchor = -1
	}
	max := len(l.rows) - 1
	start = xmath.Max(xmath.Min(start, max), 0)
	end = xmath.Max(xmath.Min(end, max), 0)
	l.Selection.SetRange(start, end)
	if l.anchor == -1 || !l.allowMultiple {
		l.anchor = start
	}
	l.MarkForRedraw()
}

// Select items at the specified indexes. If 'add' is true, then any existing selection is added to rather than
// replaced.
func (l *List[T]) Select(add bool, index ...int) {
	if !l.allowMultiple {
		add = false
		if len(index) > 0 {
			index = index[len(index)-1:]
		}
	}
	if !add {
		l.Selection.Reset()
		l.anchor = -1
	}
	max := len(l.rows)
	for _, v := range index {
		if v >= 0 && v < max {
			l.Selection.Set(v)
			if l.anchor == -1 {
				l.anchor = v
			}
		}
	}
	l.MarkForRedraw()
}

// Anchor returns the index that is the current anchor point. Will be -1 if there is no anchor point.
func (l *List[T]) Anchor() int {
	return l.anchor
}

// AllowMultipleSelection returns whether multiple rows may be selected at once.
func (l *List[T]) AllowMultipleSelection() bool {
	return l.allowMultiple
}

// SetAllowMultipleSelection sets whether multiple rows may be selected at once.
func (l *List[T]) SetAllowMultipleSelection(allow bool) *List[T] {
	l.allowMultiple = allow
	if !allow && l.Selection.Count() > 1 {
		i := l.anchor
		if i < 0 || i >= l.Count() {
			i = l.Selection.FirstSet()
		}
		l.Select(false, i)
	}
	return l
}

func (l *List[T]) rowAt(y float32) (row int, top float32) {
	count := len(l.rows)
	top = l.ContentRect(false).Y
	cellHeight := xmath.Ceil(l.Factory.CellHeight())
	if cellHeight < 1 {
		for row < count {
			_, pref, _ := l.cell(row).Sizes(Size{})
			pref.GrowToInteger()
			if top+pref.Height >= y {
				break
			}
			top += pref.Height
			row++
		}
	} else {
		row = int(xmath.Floor((y - top) / cellHeight))
		top += float32(row) * cellHeight
	}
	if row >= count {
		row = -1
		top = 0
	}
	return row, top
}

// FlashSelection flashes the current selection.
func (l *List[T]) FlashSelection() {
	l.suppressSelection = true
	l.MarkForRedraw()
	l.FlushDrawing()
	time.Sleep(l.FlashAnimationTime)
	l.suppressSelection = false
	l.MarkForRedraw()
	l.FlushDrawing()
	time.Sleep(l.FlashAnimationTime)
}
