package v1

import (
	"waf-srv/model"

	"github.com/hehanpeng/gofund/common/global/api"

	"github.com/gin-gonic/gin"
)

// 不能依赖waf 只是为了省事
func Hello(c *gin.Context) {
	var ttoInfo model.TtoInfo
	_ = c.ShouldBindJSON(&ttoInfo)
	api.Ok(c)
}
