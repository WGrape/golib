package compare

import (
	"testing"
)

// TestBiggerNum test the function BiggerNum.
func TestBiggerNum(t *testing.T) {
	if BiggerNum(4, 7) != 7 {
		t.Fail()
		return
	}

	if BiggerNum(7, 2) != 7 {
		t.Fail()
		return
	}

	if BiggerNum(7, 7) != 7 {
		t.Fail()
		return
	}
}

// TestSmallerNum test the function SmallerNum.
func TestSmallerNum(t *testing.T) {
	if SmallerNum(4, 7) != 4 {
		t.Fail()
		return
	}

	if SmallerNum(7, 2) != 2 {
		t.Fail()
		return
	}

	if SmallerNum(7, 7) != 7 {
		t.Fail()
		return
	}
}
