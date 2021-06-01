package utils

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
