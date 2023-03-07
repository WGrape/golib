// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package math provides simpler math operations.
package math

import (
	"math/rand"
	"time"
)

// GetRandInt return a rand number between left and right.
// rand number = [left, right]
func GetRandInt(left int64, right int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return left + rand.Int63n(right-left+1)
}
