package time

import (
	"github.com/WGrape/golib/array"
	"testing"
	sdtime "time"
)

var (
	testMonthList = []string{
		"2021-06", "2021-07", "2021-08", "2021-09", "2021-10",
		"2021-11", "2021-12", "2022-01", "2022-02", "2022-03",
		"2022-04", "2022-05", "2022-06", "2022-07", "2022-08",
		"2022-09", "2022-10", "2022-11", "2022-12", "2023-01",
		"2023-02", "2023-03", "2023-04", "2023-05", "2023-06",
		"2023-07", "2023-08", "2023-09", "2023-10", "2023-11",
		"2023-12", "2024-01", "2024-02", "2024-03", "2024-04",
	}
)

// TestGetNowTime test the function GetNowTime.
func TestGetNowTime(t *testing.T) {
	if GetNowTime().Format(DefaultLayout) != sdtime.Now().Format(DefaultLayout) {
		t.Fail()
	}
}

// TestGetNowTimestamp test the function GetNowTimestamp.
func TestGetNowTimestamp(t *testing.T) {
	if GetNowTimestamp() != GetNowTime().Unix() {
		t.Fail()
	}
}

// TestGetCurrentMonth tests the GetCurrentMonth function.
func TestGetCurrentMonth(t *testing.T) {
	var layout = "2006-01"
	if GetCurrentMonth(layout) != GetNowTime().Format(layout) {
		t.Fail()
		return
	}

	layout = "01"
	if GetCurrentMonth(layout) != GetNowTime().Format(layout) {
		t.Fail()
		return
	}

	layout = "200601"
	if GetCurrentMonth(layout) != GetNowTime().Format(layout) {
		t.Fail()
		return
	}
}

// TestGetCurrentDate tests the GetCurrentDate function.
func TestGetCurrentDate(t *testing.T) {
	var layout = "01-02"
	if GetCurrentDate(layout) != GetNowTime().Format(layout) {
		t.Fail()
		return
	}

	layout = "0102"
	if GetCurrentDate(layout) != GetNowTime().Format(layout) {
		t.Fail()
		return
	}

	layout = "2006-01-02"
	if GetCurrentDate(layout) != GetNowTime().Format(layout) {
		t.Fail()
		return
	}
}

// TestGetTomorrowStartTime tests the GetTomorrowStartTime function
func TestGetTomorrowStartTime(t *testing.T) {
	for _, date := range []string{"20240229", "20220720", "20220730", "20220731"} {
		startTime, err := GetTomorrowStartTime(date, "20060102")
		if err != nil {
			t.Fail()
			return
		}

		var (
			result    = startTime.Format("2006-01-02 15:04:05")
			timestamp = startTime.Unix()
		)
		if date == "20240229" && result != "2024-03-01 00:00:00" && timestamp != 1709222400 {
			t.Fail()
			return
		}
		if date == "20220720" && result != "2022-07-21 00:00:00" && timestamp != 1658332800 {
			t.Fail()
			return
		}
		if date == "20220730" && result != "2022-07-31 00:00:00" && timestamp != 1659196800 {
			t.Fail()
			return
		}
		if date == "20220731" && result != "2022-08-01 00:00:00" && timestamp != 1659283200 {
			t.Fail()
			return
		}
		t.Logf("%s, %d test success\n", result, timestamp)
	}
}

// TestGetTomorrowEndTime tests the GetTomorrowEndTime function
func TestGetTomorrowEndTime(t *testing.T) {
	for _, date := range []string{"20240229", "20220720", "20220730", "20220731"} {
		startTime, err := GetTomorrowEndTime(date, "20060102")
		if err != nil {
			t.Fail()
			return
		}

		var (
			result    = startTime.Format("2006-01-02 15:04:05")
			timestamp = startTime.Unix()
		)
		if date == "20240229" && result != "2024-03-01 23:59:59" && timestamp != 1709308799 {
			t.Fail()
			return
		}
		if date == "20220720" && result != "2022-07-21 23:59:59" && timestamp != 1658419199 {
			t.Fail()
			return
		}
		if date == "20220730" && result != "2022-07-31 23:59:59" && timestamp != 1659283199 {
			t.Fail()
			return
		}
		if date == "20220731" && result != "2022-08-01 23:59:59" && timestamp != 1659369599 {
			t.Fail()
			return
		}
		t.Logf("%s, %d test success\n", result, timestamp)
	}
}

// TestGetTomorrowDate tests the GetTomorrowDate function
func TestGetTomorrowDate(t *testing.T) {
	result, err := GetTomorrowDate("2022-04-06", DefaultDateLayout)
	if err != nil {
		t.Fail()
		return
	}

	if result != "2022-04-07" {
		t.Fail()
		return
	}
	t.Logf("%s test success\n", result)
}

// TestGetMonthStartTime tests the GetMonthStartTime function.
func TestGetMonthStartTime(t *testing.T) {
	for _, month := range testMonthList {
		startTime, err := GetMonthStartTime(month, DefaultMonthLayout)
		if err != nil {
			t.Fail()
			return
		}

		result := startTime.Format(DefaultDateLayout)
		if result != month+"-01" {
			t.Fail()
			return
		}

		t.Logf("%s test success: %s\n", month, result)
	}
}

// TestGetMonthEndTime tests the GetMonthEndTime function.
func TestGetMonthEndTime(t *testing.T) {
	// not test soFar.
	for _, month := range testMonthList {
		startTime, err := GetMonthStartTime(month, DefaultMonthLayout)
		if err != nil {
			t.Fail()
			return
		}

		endTime := GetMonthEndTime(startTime, false)

		// the 1, 3, 5, 7, 8, 10, 12 month have 31 days.
		if array.SearchInt([]int{1, 3, 5, 7, 8, 10, 12}, int(endTime.Month())) > 0 {
			result := endTime.Format(DefaultDateLayout)
			if result != month+"-31" {
				t.Fail()
				return
			}
			t.Logf("%s test success: %s\n", month, result)
		}

		// 2004 is a leap year, so February has 29 days, or February has 28 days.
		if endTime.Month() == 2 {
			result := endTime.Format(DefaultDateLayout)
			if endTime.Year() == 2024 && result != "2024-02-29" {
				t.Fail()
				return
			}
			if endTime.Year() != 2024 && result != month+"-28" {
				t.Fail()
				return
			}
			t.Logf("%s test success: %s\n", month, result)
		}

		// the 4, 6, 9, 11 month have 31 days.
		if array.SearchInt([]int{4, 6, 9, 11}, int(endTime.Month())) > 0 {
			result := endTime.Format(DefaultDateLayout)
			if result != month+"-30" {
				t.Fail()
				return
			}
			t.Logf("%s test success: %s\n", month, result)
		}
	}

	// test soFar.
	for _, month := range testMonthList {
		startTime, err := GetMonthStartTime(month, DefaultMonthLayout)
		if err != nil {
			t.Fail()
			return
		}

		endTime := GetMonthEndTime(startTime, true)
		if endTime.Unix()-GetNowTimestamp() > 86400 {
			t.Fail()
			return
		}
	}
}
