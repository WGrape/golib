package compare

import (
	"testing"
)

func TestCompare(t *testing.T) {
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
