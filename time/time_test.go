package time

import (
	"testing"
	"time"
)

// TestGetCurrentDate tests the GetCurrentDate function.
func TestGetCurrentDate(t *testing.T) {
	var layout = "01-02"
	if GetCurrentDate(layout) != time.Now().Format(layout) {
		t.Fail()
		return
	}

	layout = "0102"
	if GetCurrentDate(layout) != time.Now().Format(layout) {
		t.Fail()
		return
	}

	layout = "2006-01-02"
	if GetCurrentDate(layout) != time.Now().Format(layout) {
		t.Fail()
		return
	}
}

// TestGetCurrentMonth tests the GetCurrentMonth function.
func TestGetCurrentMonth(t *testing.T) {
	var layout = "2006-01"
	if GetCurrentMonth(layout) != time.Now().Format(layout) {
		t.Fail()
		return
	}

	layout = "01"
	if GetCurrentMonth(layout) != time.Now().Format(layout) {
		t.Fail()
		return
	}

	layout = "200601"
	if GetCurrentMonth(layout) != time.Now().Format(layout) {
		t.Fail()
		return
	}
}
