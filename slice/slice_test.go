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

func TestDelDuplicate(t *testing.T) {

	slice := []int64{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	res := DelDuplicate(slice)
	if !reflect.DeepEqual(res, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fail()
		return
	}

	s2 := []string{"aa", "bbb", "aa", "c"}
	res2 := DelDuplicate(s2)
	if !reflect.DeepEqual(res2, []string{"aa", "bbb", "c"}) {
		t.Fail()
		return
	}
}
