package slice

import (
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := Range(slice, func(n int) bool {
		return n%2 == 0
	})
	if !reflect.DeepEqual(res, []int{1, 3, 5, 7, 9}) {
		t.Fail()
		return
	}
}
