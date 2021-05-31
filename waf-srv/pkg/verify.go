package pkg

import utils "github.com/hehanpeng/gofund/common/util"

var (
	IdVerify = utils.Rules{"ID": {utils.NotEmpty()}}
)
