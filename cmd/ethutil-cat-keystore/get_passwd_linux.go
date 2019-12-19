// Copyright 2014 chaishushan@gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package main

import (
	"syscall"
	"unsafe"
)

// These constants are declared here, rather than importing
// them from the syscall package as some syscall packages, even
// on linux, for example gccgo, do not declare them.
var ioctlReadTermios = func() uintptr {
	return 0x5401 // syscall.TCGETS
}
var ioctlWriteTermios = func() uintptr {
	return 0x5402 // syscall.TCSETS
}

// terminalState contains the state of a terminal.
type terminalState struct {
	termios syscall.Termios
}

// terminalMakeRaw put the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
func terminalMakeRaw(fd int) (*terminalState, error) {
	var oldState terminalState
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlReadTermios(), uintptr(unsafe.Pointer(&oldState.termios)), 0, 0, 0); err != 0 {
		return nil, err
	}

	newState := oldState.termios
	newState.Iflag &^= syscall.ISTRIP | syscall.INLCR | syscall.ICRNL | syscall.IGNCR | syscall.IXON | syscall.IXOFF
	newState.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.ISIG
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlWriteTermios(), uintptr(unsafe.Pointer(&newState)), 0, 0, 0); err != 0 {
		return nil, err
	}

	return &oldState, nil
}

// terminalRestore restores the terminal connected to the given file descriptor to a
// previous state.
func terminalRestore(fd int, state *terminalState) error {
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlWriteTermios(), uintptr(unsafe.Pointer(&state.termios)), 0, 0, 0)
	return err
}

func getch() byte {
	if oldState, err := terminalMakeRaw(0); err != nil {
		panic(err)
	} else {
		defer terminalRestore(0, oldState)
	}

	var buf [1]byte
	if n, err := syscall.Read(0, buf[:]); n == 0 || err != nil {
		panic(err)
	}
	return buf[0]
}
