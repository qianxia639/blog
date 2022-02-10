package utils

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	timeZone   = "Asia/Shanghai"
)

// 将时间戳转换为时间格式
func TomestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// 按照指定时区解析时间格式字符串
func ParseTimeStr(timestamp time.Time) (*time.Time, error) {
	// 加载时区
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}

	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation(timeFormat, fmt.Sprintf("%v", timestamp), loc)
	if err != nil {
		return nil, err
	}

	return &timeObj, nil
}
