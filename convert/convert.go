// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package convert provides simpler conversion between int, float, and string.
package convert

import (
	"fmt"
	"strconv"
)

// ToFixedFloat64 return a float64 number with fixed decimals.
func ToFixedFloat64(value float64, format string) (float64, error) {
	if format == "" {
		format = "%.2f"
	}
	fixedFloat, err := strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	if err != nil {
		return 0, err
	}
	return fixedFloat, nil
}
