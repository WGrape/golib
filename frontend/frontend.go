// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package frontend provides an interface to work with front-end colleagues in the simpler way.
package frontend

import (
	"errors"
	"strconv"
	"strings"
)

// MakeRgbColorToInt RGB color value to three colors
func MakeRgbColorToInt(colorString string) (int, int, int, error) {
	colorString = strings.TrimPrefix(colorString, "#")
	if len(colorString) != 6 {
		return 0, 0, 0, errors.New("color length error")
	}
	color64, err := strconv.ParseInt(colorString, 16, 64)
	if err != nil {
		return 0, 0, 0, err
	}
	color := int(color64)
	return color >> 16, (color & 0x00FF00) >> 8, color & 0x0000FF, nil
}
