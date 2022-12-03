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
	"bytes"
	"encoding/binary"
	"net"
	"path/filepath"
	"time"

	"github.com/richardwilkes/json"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xio"
)

func startHandoffService(pathsChan chan<- []string, paths []string) {
	const address = "127.0.0.1:13322"
	var pathsBuffer []byte
	now := time.Now()
	for time.Since(now) < 10*time.Second {
		// First, try to establish our port and become the primary GCS instance
		if listener, err := net.Listen("tcp4", address); err == nil {
			go acceptHandoff(listener, pathsChan)
			return
		}
		if pathsBuffer == nil {
			var err error
			absPaths := make([]string, len(paths))
			for i, p := range paths {
				if absPaths[i], err = filepath.Abs(p); err != nil {
					absPaths[i] = p
				}
			}
			if pathsBuffer, err = json.Marshal(absPaths); err != nil {
				jot.Fatal(1, errs.Wrap(err))
			}
		}
		// Port is in use, try connecting as a client and handing off our file list
		if conn, err := net.DialTimeout("tcp4", address, time.Second); err == nil && handoff(conn, pathsBuffer) {
			atexit.Exit(0)
		}
		// Client can't reach the server, loop around and start the processHandoff again
	}
}

func handoff(conn net.Conn, pathsBuffer []byte) bool {
	defer xio.CloseIgnoringErrors(conn)
	buffer := make([]byte, len(cmdline.AppIdentifier))
	if n, err := conn.Read(buffer); err != nil || n != len(buffer) || !bytes.Equal(buffer, []byte(cmdline.AppIdentifier)) {
		return false
	}
	buffer = make([]byte, 5)
	buffer[0] = 22
	binary.LittleEndian.PutUint32(buffer[1:], uint32(len(pathsBuffer)))
	if n, err := conn.Write(buffer); err != nil || n != len(buffer) {
		return false
	}
	if n, err := conn.Write(pathsBuffer); err != nil || n != len(pathsBuffer) {
		return false
	}
	return true
}

func acceptHandoff(listener net.Listener, pathsChan chan<- []string) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go processHandoff(conn, pathsChan)
	}
}

func processHandoff(conn net.Conn, pathsChan chan<- []string) {
	defer xio.CloseIgnoringErrors(conn)
	if _, err := conn.Write([]byte(cmdline.AppIdentifier)); err != nil {
		return
	}
	var single [1]byte
	n, err := conn.Read(single[:])
	if err != nil {
		return
	}
	if n != 1 {
		return
	}
	if single[0] != 22 {
		return
	}
	var sizeBuffer [4]byte
	if n, err = conn.Read(sizeBuffer[:]); err != nil {
		return
	}
	if n != 4 {
		return
	}
	size := int(binary.LittleEndian.Uint32(sizeBuffer[:]))
	buffer := make([]byte, size)
	if n, err = conn.Read(buffer[:]); err != nil {
		return
	}
	if n != size {
		return
	}
	var paths []string
	if err = json.Unmarshal(buffer, &paths); err != nil {
		return
	}
	pathsChan <- paths
}
