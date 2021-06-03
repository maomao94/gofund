package router

import (
	v1 "waf-srv/api/v1"
	"waf-srv/pkg/invoker"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/server/egin"
)

func GetRouter() *egin.Component {
	Router := invoker.Gin
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("api")
	{
		PublicGroup.GET("testUps", v1.TestUps)
		PublicGroup.GET("helloLock", v1.HelloLock)
		PublicGroup.POST("hello", v1.Hello)

		// 超时转发相关
		PublicGroup.POST("registerTto", v1.RegisterTto) // 注册tto
		PublicGroup.POST("cancelTto", v1.CancelTto)     // 注销tto

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
