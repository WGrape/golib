package time

import (
	"testing"
	"time"
)

func TestGetCurrentDate(t *testing.T) {
	if GetCurrentDate() != time.Now().Format("01-02") {
		t.Fail()
		return
	}
}

func TestGetCurrentMonth(t *testing.T) {
	if GetCurrentMonth() != time.Now().Format("2006-01") {
		t.Fail()
		return
	}
}
