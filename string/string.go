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
		if !IsLetter(r) && !IsDigit(r) {
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
		if !IsLetter(r) && !IsDigit(r) {
			return false
		}

		if i == 0 && !IsLower(r) {
			return false
		} else if IsUpper(r) {
			return IsCamelCase(s[i+1:])
		}
	}
	return true
}

// IsUpper return a bool value indicating whether it is an uppercase character.
func IsUpper(r rune) bool {
	return unicode.IsUpper(r)
}

// IsLower return a bool value indicating whether it is a lowercase character.
func IsLower(r rune) bool {
	return unicode.IsLower(r)
}

// IsLetter return a bool value indicating whether it is a letter.
func IsLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

// IsDigit return a bool value indicating whether it is a digit character.
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
