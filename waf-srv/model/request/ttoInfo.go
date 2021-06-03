package request

import (
	"waf-srv/model"

	"github.com/hehanpeng/gofund/common/global/api"
)

type TtoInfoSearch struct {
	model.TtoInfo
	api.PageInfo
}

type CancelTto struct {
	ID uint `json:"id,omitempty"`
}

type RegisterTto struct {
	CallSrvName string `json:"CallSrvName"`
	CallMethod  string `json:"callMethod"`
	ExpiredTime int    `json:"expiredTime"`
}
