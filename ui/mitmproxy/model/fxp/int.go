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

package fxp

import (
	"github.com/richardwilkes/toolbox/xmath"
	"github.com/richardwilkes/toolbox/xmath/fixed"
	"github.com/richardwilkes/toolbox/xmath/fixed/f64"
)

// Common values that can be reused.
var (
	Min           = Int(f64.Min)
	NegPointEight = FromStringForced("-0.8")
	Quarter       = FromStringForced("0.25")
	Half          = FromStringForced("0.5")
	One           = From(1)
	OneAndAHalf   = FromStringForced("1.5")
	Two           = From(2)
	Three         = From(3)
	Four          = From(4)
	Five          = From(5)
	Six           = From(6)
	Seven         = From(7)
	Eight         = From(8)
	Nine          = From(9)
	Ten           = From(10)
	Twelve        = From(12)
	Fifteen       = From(15)
	Nineteen      = From(19)
	Twenty        = From(20)
	TwentyFour    = From(24)
	ThirtySix     = From(36)
	Thirty        = From(30)
	Forty         = From(40)
	Fifty         = From(50)
	Seventy       = From(70)
	Eighty        = From(80)
	NinetyNine    = From(99)
	Hundred       = From(100)
	Thousand      = From(1000)
	MaxBasePoints = From(999999)
	Max           = Int(f64.Max)
)

// DP is an alias for the fixed-point decimal places configuration we are using.
type DP = fixed.D4

// Int is an alias for the fixed-point type we are using.
type Int = f64.Int[DP]

// From creates an Int from a numeric value.
func From[T xmath.Numeric](value T) Int {
	return f64.From[DP](value)
}

// FromString creates an Int from a string.
func FromString(value string) (Int, error) {
	return f64.FromString[DP](value)
}

// FromStringForced creates an Int from a string, ignoring any conversion inaccuracies.
func FromStringForced(value string) Int {
	return f64.FromStringForced[DP](value)
}

// As returns the equivalent value in the destination type.
func As[T xmath.Numeric](value Int) T {
	return f64.As[DP, T](value)
}

// ApplyRounding rounds in the positive direction of roundDown is false, or in the negative direction if roundDown is
// true.
func ApplyRounding(value Int, roundDown bool) Int {
	if truncated := value.Trunc(); value != truncated {
		if roundDown {
			if value < 0 {
				return truncated - One
			}
		} else {
			if value > 0 {
				return truncated + One
			}
		}
		return truncated
	}
	return value
}

// ResetIfOutOfRange checks the value and if it is lower than min or greater than max, returns def, otherwise returns
// value.
func ResetIfOutOfRange(value, min, max, def Int) Int {
	if value < min || value > max {
		return def
	}
	return value
}

// ResetIfOutOfRangeInt checks the value and if it is lower than min or greater than max, returns def, otherwise returns
// value.
func ResetIfOutOfRangeInt(value, min, max, def int) int {
	if value < min || value > max {
		return def
	}
	return value
}

// Extract a leading value from a string. If a value is found, it is returned along with the portion of the string that
// was unused. If a value is not found, then 0 is returned along with the original string.
func Extract(in string) (value Int, remainder string) {
	last := 0
	max := len(in)
	if last < max && in[last] == ' ' {
		last++
	}
	if last >= max {
		return 0, in
	}
	ch := in[last]
	found := false
	decimal := false
	start := last
	for (start == last && (ch == '-' || ch == '+')) || (!decimal && ch == '.') || (ch >= '0' && ch <= '9') {
		if ch >= '0' && ch <= '9' {
			found = true
		}
		if ch == '.' {
			decimal = true
		}
		last++
		if last >= max {
			break
		}
		ch = in[last]
	}
	if !found {
		return 0, in
	}
	value, err := FromString(in[start:last])
	if err != nil {
		return 0, in
	}
	return value, in[last:]
}
