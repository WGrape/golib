// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package desensitization provides data desensitization support.
package desensitization

// DesensitizeName return the desensitized name.
func DesensitizeName(name string) string {
	var (
		runes       = []rune(name)
		length      = len(runes)
		symbol rune = '*'
	)
	if length <= 1 {
		return name
	}

	for i := 1; i <= length-1; i++ {
		runes[i] = symbol
	}
	return string(runes)
}

// DesensitizeTelNo return the desensitized telephone number.
func DesensitizeTelNo(telNo string) string {
	var (
		runes  = []rune(telNo)
		length = len(runes)
		symbol = '*'
	)
	if length != 11 {
		return telNo
	}

	for i := 3; i <= 6; i++ {
		runes[i] = symbol
	}
	return string(runes)
}

// DesensitizeIDNumber return the desensitized ID Card Number.
func DesensitizeIDNumber(telNo string) string {
	var (
		runes  = []rune(telNo)
		length = len(runes)
		symbol = '*'
	)
	if length != 18 {
		return telNo
	}

	for i := 2; i <= 15; i++ {
		runes[i] = symbol
	}
	return string(runes)
}
