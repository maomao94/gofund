package pkg

import (
	utils "github.com/hehanpeng/gofund/common/util"
)

var (
	PageInfoVerify         = utils.Rules{"Page": {utils.NotEmpty()}, "PageSize": {utils.NotEmpty()}}
	IdVerify = utils.Rules{"ID": {utils.NotEmpty()}}
)