package binary

import (
	"reflect"
	"testing"
)

// TestDecimalNumberToBinary Test the function DecimalNumberToBinary
func TestDecimalNumberToBinary(t *testing.T) {
	var intArr, strArr, str = DecimalNumberToBinary(4)
	if !reflect.DeepEqual(intArr, []int{1, 0, 0}) || !reflect.DeepEqual(strArr, []string{"1", "0", "0"}) || str != "100" {
		t.Fail()
		return
	}
}
