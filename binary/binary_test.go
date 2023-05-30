package binary

import (
	"github.com/WGrape/golib/array"
	"reflect"
	"testing"
)

// TestDecimalNumberToBinary Test the function DecimalNumberToBinary
func TestDecimalNumberToBinary(t *testing.T) {
	if !reflect.DeepEqual(array.Reverse(DecimalNumberToBinary(4)), []int{1, 0, 0}) {
		t.Fail()
		return
	}
}
