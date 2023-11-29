package slice

import (
	"fmt"
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

func TestIsExistFromSliceToT(t *testing.T) {

	fmt.Println(IsExistFromSliceToT[string]("1", []string{"2", "1", "3"})) // t
	fmt.Println(IsExistFromSliceToT[string]("1", []string{"2", "3"}))      // f
	fmt.Println(IsExistFromSliceToT[int](1, []int{2, 1, 3}))               // t

	list := []string{"apple", "banana", "orange"}
	val := "banana"
	exists := IsExistFromSliceToTV2[string](val, list, func(a, b string) bool {
		return a == b
	})
	fmt.Println(exists) // 输出: true
}

func TestFindMissingElements(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 4, 5, 6, 7}

	result := FindMissingElements(a, b)
	fmt.Println(result)
}
