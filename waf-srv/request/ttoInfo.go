package request

import (
	"waf-srv/model"

	"github.com/hehanpeng/gofund/common/global/api"
)

type TtoInfoSearch struct {
	model.TtoInfo
	api.PageInfo
}
