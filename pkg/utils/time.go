package utils

import "time"

// GetTimeMs 获取毫秒数
func GetTimeMs() int64 {
	return time.Now().UnixNano() / 1e6
}
