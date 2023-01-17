package slice

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := Range(slice, func(n int) bool {
		if n%2 != 0 {
			return false
		}
		return true
	})
	fmt.Println(res)
}
