// Copyright 2022 The golib Authors. All rights reserved.
// Use of this source code is governed by a BSD-style(https://en.wikipedia.org/wiki/BSD_licenses)
// license that can be found in the LICENSE file.

// Package time provides functionality for measuring and displaying time more efficiently.
package time

import (
	sdtime "time"
)

// Default constants
const (
	DefaultLayout      = "2006-01-02 15:04:05"
	DefaultDateLayout  = "2006-01-02"
	DefaultMonthLayout = "2006-01"
	DefaultTimeLayout  = "15:04:05"
	DefaultDayLayout   = "01-02"
)

// location variables
var (
	location = sdtime.Now().Location()
)

// GetNowTime returns current time.
func GetNowTime() sdtime.Time {
	return sdtime.Now()
}

// GetNowTimestamp returns current timestamp.
func GetNowTimestamp() int64 {
	return GetNowTime().Unix()
}

// GetCurrentMonth returns current month with different time layouts,
// such as 200601, 2006-01, 01, etc. The Default time layout is 2006-01.
func GetCurrentMonth(layout string) string {
	if layout == "" {
		layout = DefaultMonthLayout
	}
	return sdtime.Now().Format(layout)
}

// GetCurrentDate returns current date with different layouts,
// such as 01-02, 0102, 2006-01-02, etc. The Default time layout is 2006-01-02.
func GetCurrentDate(layout string) string {
	if layout == "" {
		layout = DefaultDateLayout
	}
	return sdtime.Now().Format(layout)
}

// GetDayStartTime returns the start time of day.
func GetDayStartTime(date string, layout string) (sdtime.Time, error) {
	startTime, err := sdtime.ParseInLocation(layout, date, location)
	if err != nil {
		return startTime, err
	}
	return startTime.AddDate(0, 0, 1), nil
}

// GetDayEndTime returns the end time of day.
func GetDayEndTime(date string, layout string) (sdtime.Time, error) {
	startTime, err := GetDayStartTime(date, layout)
	if err != nil {
		return startTime, err
	}
	return startTime.Add(86399 * sdtime.Second), nil
}

// GetMonthStartTime returns the start time of month.
func GetMonthStartTime(month string, layout string) (sdtime.Time, error) {
	startTime, err := sdtime.ParseInLocation(layout, month, location)
	if err != nil {
		return startTime, err
	}
	return startTime.AddDate(0, 0, -startTime.Day()+1), nil
}

// GetMonthEndTime returns the end time of month. The second param soFar is true means the return result is not greater than today.
func GetMonthEndTime(startTime sdtime.Time, soFar bool) sdtime.Time {
	endTime := sdtime.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, location).AddDate(0, 1, -1).Add(86399 * sdtime.Second)
	if soFar && endTime.Unix() >= GetNowTimestamp() {
		endTime = sdtime.Date(sdtime.Now().Year(), sdtime.Now().Month(), sdtime.Now().Day(), 0, 0, 0, 0, location).Add(86399 * sdtime.Second)
	}
	return endTime
}
