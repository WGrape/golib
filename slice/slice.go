// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package slice provides an interface to process slice in the simpler way.
package slice

import "fmt"

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

// DelDuplicate Del the duplicate elements.
func DelDuplicate(data interface{}) interface{} {
	m := make(map[string]struct{})
	switch slice := data.(type) {
	case []string:
		res := make([]string, 0, len(data.([]string)))
		for _, item := range slice {
			if _, ok := m[item]; !ok {
				m[item] = struct{}{}
				res = append(res, item)
			}
		}
		return res
	case []int64:
		res := make([]int64, 0, len(data.([]int64)))
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := m[key]; !ok {
				m[key] = struct{}{}
				res = append(res, item)
			}
		}
		return res
	default:
		return nil
	}
}
