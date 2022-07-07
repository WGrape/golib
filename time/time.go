// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package time provides functionality for measuring and displaying time more efficiently.

package time

import (
	sdtime "time"
)

// GetCurrentMonth returns current month with different layouts
func GetCurrentMonth(layout string) string {
	if layout == "" {
		layout = "2006-01"
	}
	return sdtime.Now().Format(layout)
}

// GetCurrentDate returns current date with different layouts
func GetCurrentDate(layout string) string {
	if layout == "" {
		layout = "01-02"
	}
	return sdtime.Now().Format(layout)
}
