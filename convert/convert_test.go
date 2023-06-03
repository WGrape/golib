package convert

import (
	"fmt"
	"testing"
)

// TestToFixedFloat64 test the function ToFixedFloat64
func TestToFixedFloat64(t *testing.T) {
	result, err := ToFixedFloat64(10.127, "%.2f")
	if result != 10.13 || err != nil {
		t.Fail()
		return
	}
}

// TestReverseNum test the function ReverseNum
func TestReverseNum(t *testing.T) {
	if ReverseNum(121) != 121 {
		t.Fail()
		return
	}
	fmt.Println(ReverseNum(121))

	if ReverseNum(-152) != -251 {
		t.Fail()
		return
	}
	fmt.Println(ReverseNum(-152))

	if ReverseNum(-100) != -1 {
		t.Fail()
		return
	}
	fmt.Println(ReverseNum(-100))
}
