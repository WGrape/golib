package string

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	s := "你好，世界"
	reverse := Reverse(s)
	fmt.Println(reverse) // 界世，好你

}
