// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package set provides an interface to process set in the simpler way.
package set

// Intersect return the intersection of slice1 and slice2
func Intersect(slice1, slice2 []int64) []int64 {
	var m = make(map[int64]int)
	for _, v := range slice1 {
		m[v] = 1
	}

	var result []int64
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
			delete(m, v)
		}
	}
	return result
}

// Difference return the difference of slice1 and slice2
func Difference(slice1 []int64, slice2 []int64) []int64 {
	var m = make(map[int64]int)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v] = 1
	}

	var result []int64
	for _, v := range slice1 {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = 1
		}
	}
	return result
}
