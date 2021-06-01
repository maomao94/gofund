// Package statistics 统计数据
package statistics

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
	"waf-srv/model"
	"waf-srv/pkg/invoker"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	// 输出统计数据的时间
	exportStatisticsTime = 1 * time.Second
	p                    = message.NewPrinter(language.English)
)

// ReceivingResults 接收结果并处理
// 统计的时间都是纳秒，显示的时间 都是毫秒
// concurrent 并发数
func ReceivingResults(ch <-chan *model.RequestResults, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	var stopChan = make(chan bool)
	// 时间
	var (
		processingTime uint64 // 处理总时间
		requestTime    uint64 // 请求总时间
		maxTime        uint64 // 最大时长
		minTime        uint64 // 最小时长
		successNum     uint64 // 成功处理数，code为0
		failureNum     uint64 // 处理失败数，code不为0
		chanIDLen      int    // 并发数
		chanIDs        = make(map[uint64]bool)
	)
	statTime := uint64(time.Now().UnixNano())
	// 错误码/错误个数
	var errCode = make(map[int]int)
	// 定时输出一次计算结果
	ticker := time.NewTicker(exportStatisticsTime)
	go func() {
		for {
			select {
			case <-ticker.C:
				endTime := uint64(time.Now().UnixNano())
				requestTime = endTime - statTime
				go calculateData(processingTime, requestTime, maxTime, minTime, successNum, failureNum,
					chanIDLen, errCode)
			case <-stopChan:
				// 处理完成
				return
			}
		}
	}()
	header()
	for data := range ch {
		processingTime = processingTime + data.Time
		if maxTime <= data.Time {
			maxTime = data.Time
		}
		if minTime == 0 {
			minTime = data.Time
		} else if minTime > data.Time {
			minTime = data.Time
		}
		// 是否请求成功
		if data.IsSucceed == true {
			successNum = successNum + 1
		} else {
			failureNum = failureNum + 1
		}
		// 统计错误码
		if value, ok := errCode[data.ErrCode]; ok {
			errCode[data.ErrCode] = value + 1
		} else {
			errCode[data.ErrCode] = 1
		}
		if _, ok := chanIDs[data.ChanID]; !ok {
			chanIDs[data.ChanID] = true
			chanIDLen = len(chanIDs)
		}
	}
	// 数据全部接受完成，停止定时输出统计数据
	stopChan <- true
	endTime := uint64(time.Now().UnixNano())
	requestTime = endTime - statTime
	calculateData(processingTime, requestTime, maxTime, minTime, successNum, failureNum, chanIDLen, errCode)
	invoker.Logger.Debug("*************************  结果 stat  ****************************")
	invoker.Logger.Debugf("请求总数: %d 总请求时间: %.3f秒 successNum: %d, failureNum: %d", successNum+failureNum, float64(requestTime)/1e9, successNum, failureNum)
	invoker.Logger.Debug("*************************  结果 end   ****************************")
}

// calculateData 计算数据
func calculateData(processingTime, requestTime, maxTime, minTime, successNum, failureNum uint64,
	chanIDLen int, errCode map[int]int) {
	if processingTime == 0 {
		processingTime = 1
	}
	var (
		maxTimeFloat     float64
		minTimeFloat     float64
		requestTimeFloat float64
	)
	// 纳秒=>毫秒
	maxTimeFloat = float64(maxTime) / 1e6
	minTimeFloat = float64(minTime) / 1e6
	requestTimeFloat = float64(requestTime) / 1e9
	// 打印的时长都为毫秒
	table(successNum, failureNum, errCode, maxTimeFloat, minTimeFloat, requestTimeFloat, chanIDLen)
}

// header 打印表头信息
func header() {
	// 打印的时长都为毫秒 总请数
	invoker.Logger.Debug("─────┬───────┬───────┬───────┬────────┬────────┬────────┬────────┬────────┬────────┬────────")
	invoker.Logger.Debug(" 耗时│ 并发数│ 成功数│ 失败数│最长耗时│最短耗时│ 错误码")
	invoker.Logger.Debug("─────┼───────┼───────┼───────┼────────┼────────┼────────┼────────┼────────┼────────┼────────")
	return
}

// table 打印表格
func table(successNum, failureNum uint64, errCode map[int]int, maxTimeFloat, minTimeFloat, requestTimeFloat float64, chanIDLen int) {
	// 打印的时长都为毫秒
	result := fmt.Sprintf("%4.0fs│%7d│%7d│%7d│%8.2f│%8.2f│%v",
		requestTimeFloat, chanIDLen, successNum, failureNum, maxTimeFloat, minTimeFloat, printMap(errCode))
	invoker.Logger.Debug(result)
	return
}

// printMap 输出错误码、次数 节约字符(终端一行字符大小有限)
func printMap(errCode map[int]int) (mapStr string) {
	var (
		mapArr []string
	)
	for key, value := range errCode {
		mapArr = append(mapArr, fmt.Sprintf("%d:%d", key, value))
	}
	sort.Strings(mapArr)
	mapStr = strings.Join(mapArr, ";")
	return
}
