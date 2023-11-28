// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package slice provides an interface to process slice in the simpler way.
package slice

import (
	"fmt"
	"reflect"
)

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

func FindMissingElements(a, b []int) []int {
	m := make(map[int]bool)
	var result []int

	// 将切片b中的元素存入map
	for _, num := range b {
		m[num] = true
	}

	// 遍历切片a，如果元素不在切片b中，则添加到结果切片中
	for _, num := range a {
		if _, ok := m[num]; !ok {
			result = append(result, num)
		}
	}

	return result
}

func IsExistFromSlice(val int, list []int) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

func IsExistFromSliceToString(val string, list []string) bool {
	var exist bool
	for _, v := range list {
		if val == v {
			exist = true
			break
		}
	}
	return exist
}

// IsExistFromSliceToT 泛型判断val是否存在切片中，借助反射（带来性能损失）
func IsExistFromSliceToT[T any](val T, list []T) bool {
	var exist bool
	for _, v := range list {
		if reflect.DeepEqual(v, val) {
			exist = true
			break
		}
	}
	return exist
}

type CompareFunc[T any] func(a, b T) bool

// IsExistFromSliceToTV2 泛型判断val是否存在切片中，借助自定义的比较函数（性能较高）
func IsExistFromSliceToTV2[T any](val T, list []T, compare CompareFunc[T]) bool {
	for _, v := range list {
		if compare(val, v) {
			return true
		}
	}
	return false
}
