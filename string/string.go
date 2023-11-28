// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package string provides an interface to process string in the simpler way.
package string

import "unicode"

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsValidIdentifierName return a bool value indicating whether it is a valid identifier name.
func IsValidIdentifierName(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !IsLetter([]rune{r}) && !IsDigit([]rune{r}) {
			return false
		}
	}
	return true
}

// IsCamelCase return a bool value indicating whether it is a valid camel case spelling.
func IsCamelCase(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i, r := range s {
		if !IsLetter([]rune{r}) && !IsDigit([]rune{r}) {
			return false
		}

		if i == 0 && !IsLower([]rune{r}) {
			return false
		} else if IsUpper([]rune{r}) {
			return IsCamelCase(s[i+1:])
		}
	}
	return true
}

// IsUpper return a bool value indicating whether it is an uppercase character.
func IsUpper(str []rune) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// IsLower return a bool value indicating whether it is a lowercase character.
func IsLower(str []rune) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

// IsLetter return a bool value indicating whether it is a letter.
func IsLetter(str []rune) bool {
	for _, r := range str {
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			return false
		}
	}
	return true
}

// IsDigit return a bool value indicating whether it is a digit character.
func IsDigit(str []rune) bool {
	for _, r := range str {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
