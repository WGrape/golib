// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package binary provides some operations of binary number and decimal number.
package binary

import "fmt"

// DecimalNumberToBinary Return a binary array.
func DecimalNumberToBinary(num int) ([]int, []string, string) {
	var (
		binary       []int
		binaryString []string
		str          string
	)

	for num != 0 {
		var (
			bit    = num % 2
			bitStr = fmt.Sprintf("%d", bit)
		)
		binary = append([]int{bit}, binary...)
		binaryString = append([]string{bitStr}, binaryString...)
		str = bitStr + str
		num = num / 2
	}

	return binary, binaryString, str
}
