package core

import (
	"github.com/gin-gonic/gin"
)

// HandlerFunc core封装后的handler
type HandlerFunc func(c *Context)

// Handle 将core.HandlerFunc转换为gin.HandlerFunc
func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

// Context core封装后的Context
type Context struct {
	*gin.Context
}
