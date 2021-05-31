package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/server/egin"
	v1 "waf-srv/api/v1"
	"waf-srv/pkg/invoker"
)

func GetRouter() *egin.Component {
	Router := invoker.Gin
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("api")
	{
		InitTtoInfoRouter(PublicGroup)
	}
	return Router
}

func InitTtoInfoRouter(Router *gin.RouterGroup) {
	TtoInfoRouter := Router.Group("ttoInfo")
	{
		TtoInfoRouter.POST("createTtoInfo", v1.CreateTtoInfo)             // 新建TtoInfo
		TtoInfoRouter.DELETE("deleteTtoInfo", v1.DeleteTtoInfo)           // 删除TtoInfo
		TtoInfoRouter.DELETE("deleteTtoInfoByIds", v1.DeleteTtoInfoByIds) // 批量删除TtoInfo
		TtoInfoRouter.PUT("updateTtoInfo", v1.UpdateTtoInfo)              // 更新TtoInfo
		TtoInfoRouter.GET("findTtoInfo", v1.FindTtoInfo)                  // 根据ID获取TtoInfo
		TtoInfoRouter.GET("getTtoInfoList", v1.GetTtoInfoList)            // 获取TtoInfo列表
	}
}
