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

// GetTomorrowStartTime returns the start time of tomorrow.
func GetTomorrowStartTime(date string, layout string) (sdtime.Time, error) {
	startTime, err := sdtime.ParseInLocation(layout, date, location)
	if err != nil {
		return startTime, err
	}
	return startTime.AddDate(0, 0, 1), nil
}

// GetTomorrowEndTime returns the end time of tomorrow.
func GetTomorrowEndTime(date string, layout string) (sdtime.Time, error) {
	startTime, err := GetTomorrowStartTime(date, layout)
	if err != nil {
		return startTime, err
	}
	return startTime.Add(86399 * sdtime.Second), nil
}

// GetTomorrowDate returns tomorrow date.
func GetTomorrowDate(date string, layout string) (string, error) {
	tomorrow, err := GetTomorrowStartTime(date, layout)
	if err != nil {
		return "", err
	}

	return tomorrow.Format(layout), nil
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

// ParseIso8601ToTime converts the iso8601 time to time type, about iso8601 is here https://zh.m.wikipedia.org/zh-hans/ISO_8601
func ParseIso8601ToTime(iso8601TimeString string) (sdtime.Time, error) {
	result, err := sdtime.ParseInLocation("2006-01-02T15:04:05+08:00", iso8601TimeString, sdtime.Local)
	if err != nil {
		return sdtime.Time{}, err
	}
	return result, nil
}
