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
	Reference   int    `json:"reference"`
	TtoType     string `json:"tto_type"`
	BizType     string `json:"biz_type"`
	CallSrvName string `json:"CallSrvName"`
	CallMethod  string `json:"callMethod"`
	ExpiredTime int    `json:"expiredTime"`
	Ext1        string `json:"ext_1"`
}
