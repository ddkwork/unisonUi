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

import "github.com/richardwilkes/unison"

// NewInteriorSeparator creates a new interior vertical separator.
func NewInteriorSeparator() *unison.Separator {
	spacer := unison.NewSeparator()
	spacer.LineInk = unison.InteriorDividerColor
	spacer.Vertical = true
	spacer.SetLayoutData(&unison.FlexLayoutData{VAlign: unison.FillAlignment})
	return spacer
}

// NewToolbarSeparator creates a new vertical separator for the toolbar.
func NewToolbarSeparator() *unison.Separator {
	spacer := unison.NewSeparator()
	spacer.LineInk = unison.ControlEdgeColor
	spacer.Vertical = true
	spacer.SetBorder(unison.NewEmptyBorder(unison.NewHorizontalInsets(unison.StdHSpacing)))
	spacer.SetLayoutData(&unison.FlexLayoutData{VAlign: unison.FillAlignment})
	spacer.DrawCallback = func(canvas *unison.Canvas, _ unison.Rect) {
		rect := spacer.ContentRect(false)
		paint := spacer.LineInk.Paint(canvas, rect, unison.Stroke)
		paint.SetPathEffect(unison.NewDashPathEffect([]float32{2, 2}, 0))
		canvas.DrawLine(rect.X, rect.Y, rect.X, rect.Bottom(), paint)
	}
	return spacer
}
