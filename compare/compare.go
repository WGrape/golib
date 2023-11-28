// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package compare provides simpler comparison.
package compare

// BiggerNum Return a bigger number
func BiggerNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SmallerNum Return a smaller number
func SmallerNum(a, b int) int {
	if a > b {
		return b
	}
	return a
}
