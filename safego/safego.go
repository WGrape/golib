// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package safego provides a set of security behaviors.
package safego

import (
	"os"
	"os/signal"
	"syscall"
)

// ListenServerStop listen the server is stopped
func ListenServerStop(isStopped *bool) <-chan bool {
	var (
		wait    = make(chan bool)
		sigChan = make(chan os.Signal, 1)
	)
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		s := <-sigChan
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			*isStopped = true
			close(wait)
			return wait
		}
	}
}
