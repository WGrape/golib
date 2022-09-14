package redis

import "github.com/WGrape/golib/math"

// GetStringEmptyObject return empty object of string for writing it to cache
func GetStringEmptyObject() string {
	return ""
}

// GetExpireSeconds get expire seconds with random
func GetExpireSeconds(seconds int64) int64 {
	return seconds + math.GetRandInt(1, 100)
}
