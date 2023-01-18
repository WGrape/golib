package string

import (
	"testing"
)

func TestReverse(t *testing.T) {
	s := "hello world"
	if Reverse(s) != "dlrow olleh" {
		t.Fail()
		return
	}
}
