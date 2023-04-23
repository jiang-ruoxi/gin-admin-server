package utils

import (
	"fmt"
	"time"
)

const (
	timeFormatTpl = "2006-01-02"
	tf2           = "2006-01-02 15:04:05"
	RFC3339       = "2006-01-02T15:04:05Z07:00"
)

// 时间戳转格式化
func FormatDateFromUnix(unix int64) string {
	date := time.Unix(unix, 0)
	return date.Format(tf2)
}

// 日期只是具体的年月日，没有时分秒
func FormatDateToUnixTime(d string) int64 {
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.Parse(RFC3339, d)
	if err != nil {
		fmt.Println("Failed to parse date string:", err)
		return 0
	}
	date := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, TimeLocation)
	return date.Unix()
}