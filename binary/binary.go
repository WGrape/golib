// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package binary provides some operations of binary number and decimal number.
package binary

// DecimalNumberToBinary Return a binary array.
func DecimalNumberToBinary(num int) []int {
	var binary []int

	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}

	return binary
}
