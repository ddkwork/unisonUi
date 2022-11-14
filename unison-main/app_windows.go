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
	"sync/atomic"
	"syscall"
	"time"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xio"
	"github.com/richardwilkes/unison/internal/w32"
	"golang.org/x/sys/windows/registry"
)

var appUsesLightThemeValue = uint32(1)

func platformEarlyInit() {
	AttachConsole()
}

func platformLateInit() {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`, syscall.KEY_NOTIFY|registry.QUERY_VALUE)
	if err != nil {
		jot.Warn(errs.NewWithCause("unable to open dark mode key", err))
		return
	}
	if err = updateTheme(k, true); err != nil {
		jot.Warn(err)
		xio.CloseIgnoringErrors(k)
		return
	}
	go func() {
		for {
			w32.RegNotifyChangeKeyValue(k, false, w32.RegNotifyChangeName|w32.RegNotifyChangeLastSet, 0, false)
			if err = updateTheme(k, false); err != nil {
				jot.Warn(err)
				xio.CloseIgnoringErrors(k)
				return
			}
		}
	}()
}

func platformBeep() {
	w32.MessageBeep(w32.MBDefault)
}

func platformIsDarkModeTrackingPossible() bool {
	return true
}

func platformIsDarkModeEnabled() bool {
	return atomic.LoadUint32(&appUsesLightThemeValue) == 0
}

func platformDoubleClickInterval() time.Duration {
	return w32.GetDoubleClickTime()
}

func updateTheme(k registry.Key, sync bool) error {
	val, _, err := k.GetIntegerValue("AppsUseLightTheme")
	if err != nil {
		return errs.NewWithCause("unable to retrieve dark mode key", err)
	}
	var swapped bool
	if val == 0 {
		swapped = atomic.CompareAndSwapUint32(&appUsesLightThemeValue, 1, 0)
	} else {
		swapped = atomic.CompareAndSwapUint32(&appUsesLightThemeValue, 0, 1)
	}
	if swapped && currentColorMode == AutomaticColorMode {
		if sync {
			ThemeChanged()
		} else {
			InvokeTask(ThemeChanged)
		}
	}
	return nil
}
