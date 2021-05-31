package request

import (
	"github.com/hehanpeng/gofund/common/req"
	"waf-srv/model"
)

type TtoInfoSearch struct {
	model.TtoInfo
	req.PageInfo
}
