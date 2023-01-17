// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package slice provides an interface to process slice in the simpler way.
package slice

// Range range a slice and return a new slice after deleting target items.
func Range(data []int, f func(n int) bool) []int {
	k := 0
	res := make([]int, len(data))
	for _, item := range data {
		if !f(item) {
			res[k] = item
			k++
		}
	}
	res = res[:k]
	return res
}
