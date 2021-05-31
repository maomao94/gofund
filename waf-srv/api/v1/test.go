package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hehanpeng/gofund/common/response"
)

func Test(c *gin.Context) {
	response.Ok(c)
}
