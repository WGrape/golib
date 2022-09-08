package math

import "testing"

func TestGetRandInt(t *testing.T) {
	for i := 1; i <= 200; i++ {
		if GetRandInt(0, 0) != 0 {
			t.Fail()
			return
		}
	}

	for i := 1; i <= 200; i++ {
		if GetRandInt(7, 7) != 7 {
			t.Fail()
			return
		}
	}

	for i := 1; i <= 200; i++ {
		var v = GetRandInt(10, 11)
		if v > 11 || v < 10 {
			t.Fail()
			return
		}
	}

	for i := 1; i <= 200; i++ {
		var v = GetRandInt(10, 20)
		if v > 20 || v < 10 {
			t.Fail()
			return
		}
	}
}
