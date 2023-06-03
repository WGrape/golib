// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package rand provides some random operations.
package rand

import (
	crand "crypto/rand"
	"encoding/hex"
	"io"
)

// GenUUid Generate the uuid.
func GenUUid(bufBytes int) string {
	var buf = make([]byte, bufBytes)
	_, _ = io.ReadFull(crand.Reader, buf)
	var uuid = hex.EncodeToString(buf)
	return uuid
}
