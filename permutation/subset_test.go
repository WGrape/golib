package permutation

import (
	"testing"
)

func TestSubsets(t *testing.T) {
	if len(Subsets([]int{0, 1, 3, 5, 9})) != 32 {
		t.Fail()
		return
	}
}
