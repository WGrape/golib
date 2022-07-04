package time

import (
	sdtime "time"
)

// GetCurrentMonth 获取当前的月份
func GetCurrentMonth() string {
	return sdtime.Now().Format("2006-01")
}

// GetCurrentDate 获取当前的日期
func GetCurrentDate() string {
	return sdtime.Now().Format("01-02")
}
