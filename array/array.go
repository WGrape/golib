// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package array provides functionality for measuring and displaying array more efficiently.
package array

// InStrings return bool means the value is exists in the target string array or not.
func InStrings(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// InInts return bool means the value is exists in the target int array or not.
func InInts(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// IsStringsEqual return bool means the two arrays are equal or not.
func IsStringsEqual(arr1 []string, arr2 []string) bool {
	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
