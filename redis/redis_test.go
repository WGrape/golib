package redis

import "testing"

// TestGetStringEmptyObject test the function GetStringEmptyObject
func TestGetStringEmptyObject(t *testing.T) {
	if GetStringEmptyObject() != "" {
		t.Fail()
		return
	}
}

// TestGetExpireSeconds test the function GetExpireSeconds
func TestGetExpireSeconds(t *testing.T) {
	var expireSeconds int64 = 1000
	for i := 1; i <= 100; i++ {
		if GetExpireSeconds(expireSeconds) <= expireSeconds {
			t.Fail()
			return
		}
	}
}
