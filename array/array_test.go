package array

import "testing"

// TestSearchString test the function SearchString
func TestSearchString(t *testing.T) {
	if SearchString([]string{"123", "456", "789"}, "456") != 1 {
		t.Fail()
		return
	}
}

// TestSearchInt test the function SearchInt
func TestSearchInt(t *testing.T) {
	if SearchInt([]int{123, 456, 789}, int(456)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchUInt test the function SearchUInt
func TestSearchUInt(t *testing.T) {
	if SearchUInt([]uint{123, 456, 789}, uint(456)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchInt64 test the function SearchInt64
func TestSearchInt64(t *testing.T) {
	if SearchInt64([]int64{123, 456, 789}, int64(456)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchUInt64 test the function SearchUInt64
func TestSearchUInt64(t *testing.T) {
	if SearchUInt64([]uint64{123, 456, 789}, uint64(456)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchInt32 test the function SearchInt32
func TestSearchInt32(t *testing.T) {
	if SearchInt32([]int32{123, 456, 789}, int32(456)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchUInt32 test the function SearchUInt32
func TestSearchUInt32(t *testing.T) {
	if SearchUInt32([]uint32{123, 456, 789}, uint32(456)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchInt16 test the function SearchInt16
func TestSearchInt16(t *testing.T) {
	if SearchInt16([]int16{123, 456, 789}, int16(456)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchUInt16 test the function SearchUInt16
func TestSearchUInt16(t *testing.T) {
	if SearchUInt16([]uint16{123, 456, 789}, uint16(456)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchInt8 test the function SearchInt8
func TestSearchInt8(t *testing.T) {
	if SearchInt8([]int8{12, 34, 56}, int8(34)) != 1 {
		t.Fail()
		return
	}
}

// TestSearchUInt8 test the function SearchUInt8
func TestSearchUInt8(t *testing.T) {
	if SearchUInt8([]uint8{12, 34, 56}, uint8(34)) != 1 {
		t.Fail()
		return
	}
}

func TestIsStringsEqual(t *testing.T) {
	if IsStringsEqual([]string{"123"}, []string{"456"}) {
		t.Fail()
		return
	}
	if !IsStringsEqual([]string{"123"}, []string{"123"}) {
		t.Fail()
		return
	}
}
