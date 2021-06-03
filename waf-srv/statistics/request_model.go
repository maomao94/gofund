package statistics

import "fmt"

// RequestResults 请求结果
type RequestResults struct {
	ID        string // 消息ID
	ChanID    uint64 // 消息ID
	Time      uint64 // 请求时间 纳秒
	IsSucceed bool   // 是否请求成功
	ErrCode   int    // 错误码
}

// SetID 设置请求唯一ID
func (r *RequestResults) SetID(chanID uint64, number uint64) {
	id := fmt.Sprintf("%d_%d", chanID, number)
	r.ID = id
	r.ChanID = chanID
}
