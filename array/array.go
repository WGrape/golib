// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package array provides functionality for measuring and displaying array more efficiently.
package array

import (
	"math/rand"
	"time"
)

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

// ShuffleIntSlice shuffle an int slice randomly.
func ShuffleIntSlice(s []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleUIntSlice shuffle an uint slice randomly.
func ShuffleUIntSlice(s []uint) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleInt64Slice shuffle an int64 slice randomly.
func ShuffleInt64Slice(s []int64) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleUInt64Slice shuffle an uint64 slice randomly.
func ShuffleUInt64Slice(s []uint64) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleInt32Slice shuffle an int32 slice randomly.
func ShuffleInt32Slice(s []int32) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleUInt32Slice shuffle an uint32 slice randomly.
func ShuffleUInt32Slice(s []uint32) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleInt16Slice shuffle an int16 slice randomly.
func ShuffleInt16Slice(s []int16) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleUInt16Slice shuffle an uint16 slice randomly.
func ShuffleUInt16Slice(s []uint16) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleInt8Slice shuffle an int8 slice randomly.
func ShuffleInt8Slice(s []int8) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// ShuffleUInt8Slice shuffle an uint8 slice randomly.
func ShuffleUInt8Slice(s []uint8) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
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
