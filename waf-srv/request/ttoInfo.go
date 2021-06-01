package request

import (
	"waf-srv/model"

	"github.com/hehanpeng/gofund/common/global"
)

type TtoInfoSearch struct {
	model.TtoInfo
	global.PageInfo
}
