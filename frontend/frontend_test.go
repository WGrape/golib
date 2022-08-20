package frontend

import "testing"

// TestMakeRgbColorToInt test the function of MakeRgbColorToInt
func TestMakeRgbColorToInt(t *testing.T) {
	a, b, c, _ := MakeRgbColorToInt("#0864c2")
	if a != 8 || b != 100 || c != 194 {
		t.Fail()
		return
	}

	a, b, c, _ = MakeRgbColorToInt("#0865c3")
	if a != 8 || b != 101 || c != 195 {
		t.Fail()
		return
	}

	a, b, c, _ = MakeRgbColorToInt("#102cca")
	if a != 16 || b != 44 || c != 202 {
		t.Fail()
		return
	}

	a, b, c, _ = MakeRgbColorToInt("#000000")
	if a != 0 || b != 0 || c != 0 {
		t.Fail()
		return
	}

	a, b, c, _ = MakeRgbColorToInt("#ffffff")
	if a != 255 || b != 255 || c != 255 {
		t.Fail()
		return
	}

	a, b, c, _ = MakeRgbColorToInt("#888888")
	if a != 136 || b != 136 || c != 136 {
		t.Fail()
		return
	}

	_, _, _, err := MakeRgbColorToInt("#888")
	if err == nil {
		t.Fail()
		return
	}
}
