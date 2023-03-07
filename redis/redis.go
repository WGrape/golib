// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package redis provides advanced operations, such as Cache Penetration, Cache Breakdown, Cache Avalanche.
package redis

import "github.com/WGrape/golib/math"

// GetStringEmptyObject return empty object of string for writing it to cache
func GetStringEmptyObject() string {
	return ""
}

// GetExpireSeconds get expire seconds with random
func GetExpireSeconds(seconds int64) int64 {
	return seconds + math.GetRandInt(1, 100)
}
