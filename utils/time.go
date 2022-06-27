package utils

import (
	"fmt"
	"time"
)

// 将时间戳秒数转换为时间格式
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// 将时间戳转换为 YYYY-MM-dd 格式
func TimestampToString(timestamp int64) string {
	year, month, day := TimestampToTime(timestamp).Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
