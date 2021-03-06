// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package permutation provides many algorithms about permutation and combination.
package permutation

import (
	"strings"
)

// generate will find out all combinations and fill them into the resultList
func generate(start int, targetLen int, nums []string, tmpArray []string) [][]string {
	if len(tmpArray) == targetLen {
		part := make([]string, len(tmpArray))
		copy(part, tmpArray)
		return [][]string{part}
	}

	var resultList [][]string
	for i := start; i <= len(nums)-1; i++ {
		tmpArray = append(tmpArray, nums[i])
		resultList = append(resultList, generate(i+1, targetLen, nums, tmpArray)...)
		tmpArray = tmpArray[:len(tmpArray)-1]
	}

	return resultList
}

// GetCombinations return the all combinations of the given string array.
func GetCombinations(array []string) [][]string {
	var resultList [][]string
	for targetLen := len(array); targetLen >= 1; targetLen-- {
		resultList = append(resultList, generate(0, targetLen, array, make([]string, 0))...)
	}
	return resultList
}

// GetCombinationsWithImplode return the all combinations by implode every combination.
func GetCombinationsWithImplode(nums []string, sep string) []string {
	if sep == "" {
		sep = ";"
	}
	tmpList := GetCombinations(nums)
	result := make([]string, 0)
	for _, v := range tmpList {
		result = append(result, strings.Join(v, sep))
	}
	return result
}
