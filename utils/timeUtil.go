package utils

import (
	"time"
)

// 将时间戳转换为时间格式
func TimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
