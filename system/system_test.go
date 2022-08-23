package system

import "testing"

// TestGetGoPath test the function GetGoPath
func TestGetGoPath(t *testing.T) {
	if GetGoPath() == "" {
		t.Fail()
		return
	}
}
