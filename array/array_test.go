package array

import (
	"fmt"
	"reflect"
	"testing"
)

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

// TestIsStringsEqual test the function IsStringsEqual
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

// TestShuffleIntSlice test the function ShuffleIntSlice
func TestShuffleIntSlice(t *testing.T) {
	var s = []int{1, 9, 6, 3, 5, 2}
	ShuffleIntSlice(s)
	fmt.Println(s)
}

// TestShuffleUIntSlice test the function ShuffleUIntSlice
func TestShuffleUIntSlice(t *testing.T) {
	var s = []uint{1, 9, 6, 3, 5, 2}
	ShuffleUIntSlice(s)
	fmt.Println(s)
}

// TestShuffleInt64Slice test the function ShuffleInt64Slice
func TestShuffleInt64Slice(t *testing.T) {
	var s = []int64{1, 9, 6, 3, 5, 2}
	ShuffleInt64Slice(s)
	fmt.Println(s)
}

// TestShuffleUInt64Slice test the function ShuffleUInt64Slice
func TestShuffleUInt64Slice(t *testing.T) {
	var s = []uint64{1, 9, 6, 3, 5, 2}
	ShuffleUInt64Slice(s)
	fmt.Println(s)
}

// TestShuffleInt32Slice test the function ShuffleInt32Slice
func TestShuffleInt32Slice(t *testing.T) {
	var s = []int32{1, 9, 6, 3, 5, 2}
	ShuffleInt32Slice(s)
	fmt.Println(s)
}

// TestShuffleUInt32Slice test the function ShuffleUInt32Slice
func TestShuffleUInt32Slice(t *testing.T) {
	var s = []uint32{1, 9, 6, 3, 5, 2}
	ShuffleUInt32Slice(s)
	fmt.Println(s)
}

// TestShuffleInt16Slice test the function ShuffleInt16Slice
func TestShuffleInt16Slice(t *testing.T) {
	var s = []int16{1, 9, 6, 3, 5, 2}
	ShuffleInt16Slice(s)
	fmt.Println(s)
}

// TestShuffleUInt16Slice test the function ShuffleUInt16Slice
func TestShuffleUInt16Slice(t *testing.T) {
	var s = []uint16{1, 9, 6, 3, 5, 2}
	ShuffleUInt16Slice(s)
	fmt.Println(s)
}

// TestShuffleInt8Slice test the function ShuffleInt8Slice
func TestShuffleInt8Slice(t *testing.T) {
	var s = []int8{1, 9, 6, 3, 5, 2}
	ShuffleInt8Slice(s)
	fmt.Println(s)
}

// TestShuffleUInt8Slice test the function ShuffleUInt8Slice
func TestShuffleUInt8Slice(t *testing.T) {
	var s = []uint8{1, 9, 6, 3, 5, 2}
	ShuffleUInt8Slice(s)
	fmt.Println(s)
}

// TestPopOneInt Test the function PopOneInt
func TestPopOneInt(t *testing.T) {
	var (
		e    int
		list = []int{1, 3, 4, 6, 7}
	)
	if e, list = PopOneInt(list); e != 7 {
		t.Fail()
		return
	}
	if !reflect.DeepEqual(list, []int{1, 3, 4, 6}) {
		t.Fail()
		return
	}
}

// TestReverse Test the function Reverse
func TestReverse(t *testing.T) {
	if !reflect.DeepEqual(Reverse([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}) {
		t.Fail()
		return
	}
}
