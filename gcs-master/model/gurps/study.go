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

	"github.com/richardwilkes/gcs/v5/model/fxp"
	"github.com/richardwilkes/toolbox/i18n"
)

// Study holds data about a single study session.
type Study struct {
	Type  StudyType `json:"type"`
	Hours fxp.Int   `json:"hours"`
	Note  string    `json:"note,omitempty"`
}

// Clone creates a copy of the TemplatePicker.
func (s *Study) Clone() *Study {
	clone := *s
	return &clone
}

// ResolveStudyHours returns the resolved total study hours.
func ResolveStudyHours(study []*Study) fxp.Int {
	var total fxp.Int
	for _, one := range study {
		total += one.Hours.Mul(one.Type.Multiplier())
	}
	return total
}

// StudyHoursProgressText returns the progress text or an empty string.
func StudyHoursProgressText(hours fxp.Int) string {
	if hours <= 0 {
		return ""
	}
	return fmt.Sprintf(i18n.Text("Studied %v/200 hours"), hours)
}
