package helper

import (
	"fmt"
	"time"
)

func TimeCost(start time.Time) {
	tc := time.Since(start)
	fmt.Printf("time cost = %v\n", tc)
}

func MillisecondCost(startTime uint64) float64 {
	endTime := uint64(time.Now().UnixNano())
	// 纳秒=>毫秒
	return float64(endTime-startTime) / 1e9
}

// DiffNano 时间差，纳秒
func DiffNano(startTime time.Time) (diff int64) {
	diff = int64(time.Since(startTime))
	return
}

// InArrayStr 判断字符串是否在数组内
func InArrayStr(str string, arr []string) (inArray bool) {
	for _, s := range arr {
		if s == str {
			inArray = true
			break
		}
	}
	return
}
