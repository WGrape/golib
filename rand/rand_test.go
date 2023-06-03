package rand

import (
	"testing"
)

// TestGenUUid Test the function GenUUid
func TestGenUUid(t *testing.T) {
	var m = make(map[string]bool)
	for i := 1; i <= 9999999; i++ {
		var uuid = GenUUid(16)
		if m[uuid] {
			t.Fail()
			return
		}
		m[uuid] = true
	}
}
