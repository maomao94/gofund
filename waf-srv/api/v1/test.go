package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hehanpeng/gofund/common/resp"
)

func Test(c *gin.Context) {
	resp.Ok(c)
}
