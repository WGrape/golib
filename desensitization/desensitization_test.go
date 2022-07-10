package desensitization

import "testing"

// TestDesensitizeName test the function DesensitizeName.
func TestDesensitizeName(t *testing.T) {
	var result = DesensitizeName("王五")
	if result != "王*" {
		t.Fail()
		return
	}
	t.Logf("%s test sucess", result)
}

// TestDesensitizeTelNo test the function DesensitizeTelNo.
func TestDesensitizeTelNo(t *testing.T) {
	var result = DesensitizeTelNo("12231230000")
	if result != "122****0000" {
		t.Fail()
		return
	}
	t.Logf("%s test sucess", result)
}

// TestDesensitizeIDNumber test the function DesensitizeIDNumber.
func TestDesensitizeIDNumber(t *testing.T) {
	var result = DesensitizeIDNumber("410001188888880000")
	if result != "41**************00" {
		t.Fail()
		return
	}
	t.Logf("%s test sucess", result)
}
