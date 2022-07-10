// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package array provides functionality for measuring and displaying array more efficiently.
package array

// SearchString search the value in the string array and return index or return -1.
func SearchString(arr []string, value string) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchInt search the value in the int array and return index or return -1.
func SearchInt(arr []int, value int) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchUInt search the value in the int array and return index or return -1.
func SearchUInt(arr []uint, value uint) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchInt64 search the value in the int64 array and return index or return -1.
func SearchInt64(arr []int64, value int64) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchUInt64 search the value in the uint64 array and return index or return -1.
func SearchUInt64(arr []uint64, value uint64) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchInt32 search the value in the int32 array and return index or return -1.
func SearchInt32(arr []int32, value int32) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchUInt32 search the value in the uint32 array and return index or return -1.
func SearchUInt32(arr []uint32, value uint32) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchInt16 search the value in the int16 array and return index or return -1.
func SearchInt16(arr []int16, value int16) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchUInt16 search the value in the uint16 array and return index or return -1.
func SearchUInt16(arr []uint16, value uint16) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchInt8 search the value in the int8 array and return index or return -1.
func SearchInt8(arr []int8, value int8) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
}

// SearchUInt8 search the value in the uint8 array and return index or return -1.
func SearchUInt8(arr []uint8, value uint8) int {
	for index, v := range arr {
		if v == value {
			return index
		}
	}
	return -1
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
