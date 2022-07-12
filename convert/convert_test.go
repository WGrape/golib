package convert

import "testing"

// TestToFixedFloat64 test the function ToFixedFloat64
func TestToFixedFloat64(t *testing.T) {
	result, err := ToFixedFloat64(10.127, "%.2f")
	if result != 10.13 || err != nil {
		t.Fail()
		return
	}
}
