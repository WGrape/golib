package binary

import (
	"reflect"
	"testing"
)

// TestDecimalNumberToBinary Test the function DecimalNumberToBinary
func TestDecimalNumberToBinary(t *testing.T) {
	if reflect.DeepEqual(DecimalNumberToBinary(4), []int{1, 0, 0}) {
		t.Fail()
		return
	}
}
