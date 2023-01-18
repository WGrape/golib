package slice

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := Range(slice, func(n int) bool {
		if n%2 == 0 {
			return true
		}
		return false
	})
	fmt.Println(res)

}

func TestDelDuplicate(t *testing.T) {

	slice := []int64{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	res := DelDuplicate(slice)
	fmt.Println(res)
	fmt.Println("-------------------")

	s2 := []string{"aa", "bbb", "aa", "c"}
	res2 := DelDuplicate(s2)
	fmt.Println(res2)

}
