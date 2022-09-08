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
