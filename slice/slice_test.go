package slice

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := Range[int](slice, func(n int) bool {
		if n%2 == 0 {
			return true
		}
		return false
	})
	fmt.Println(res)
	fmt.Println("----------------")

	s2 := []string{"aa", "bb", "cc", "ddd"}
	res2 := Range[string](s2, func(n string) bool {
		if len(n) == 2 {
			return true
		}
		return false
	})
	fmt.Println(res2)
}
