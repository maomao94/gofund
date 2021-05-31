package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hehanpeng/gofund/common/req"
	"github.com/hehanpeng/gofund/common/resp"
	"go.uber.org/zap"
	"waf-srv/model"
	"waf-srv/pkg/invoker"
	"waf-srv/request"
	"waf-srv/service"
)

// @Tags TtoInfo
// @Summary 创建TtoInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TtoInfo true "创建TtoInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ttoInfo/createTtoInfo [post]
func CreateTtoInfo(c *gin.Context) {
	var ttoInfo model.TtoInfo
	_ = c.ShouldBindJSON(&ttoInfo)
	if err := service.CreateTtoInfo(ttoInfo); err != nil {
		invoker.Logger.Error("创建失败!", zap.Any("err", err))
		resp.FailWithMessage("创建失败", c)
	} else {
		resp.OkWithMessage("创建成功", c)
	}
}

// @Tags TtoInfo
// @Summary 删除TtoInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TtoInfo true "删除TtoInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ttoInfo/deleteTtoInfo [delete]
func DeleteTtoInfo(c *gin.Context) {
	var ttoInfo model.TtoInfo
	_ = c.ShouldBindJSON(&ttoInfo)
	if err := service.DeleteTtoInfo(ttoInfo); err != nil {
		invoker.Logger.Error("删除失败!", zap.Any("err", err))
		resp.FailWithMessage("删除失败", c)
	} else {
		resp.OkWithMessage("删除成功", c)
	}
}

// @Tags TtoInfo
// @Summary 批量删除TtoInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body req.IdsReq true "批量删除TtoInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ttoInfo/deleteTtoInfoByIds [delete]
func DeleteTtoInfoByIds(c *gin.Context) {
	var IDS req.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTtoInfoByIds(IDS); err != nil {
		invoker.Logger.Error("批量删除失败!", zap.Any("err", err))
		resp.FailWithMessage("批量删除失败", c)
	} else {
		resp.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TtoInfo
// @Summary 更新TtoInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TtoInfo true "更新TtoInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ttoInfo/updateTtoInfo [put]
func UpdateTtoInfo(c *gin.Context) {
	var ttoInfo model.TtoInfo
	_ = c.ShouldBindJSON(&ttoInfo)
	if err := service.UpdateTtoInfo(ttoInfo); err != nil {
		invoker.Logger.Error("更新失败!", zap.Any("err", err))
		resp.FailWithMessage("更新失败", c)
	} else {
		resp.OkWithMessage("更新成功", c)
	}
}

// @Tags TtoInfo
// @Summary 用id查询TtoInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TtoInfo true "用id查询TtoInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ttoInfo/findTtoInfo [get]
func FindTtoInfo(c *gin.Context) {
	var ttoInfo model.TtoInfo
	_ = c.ShouldBindQuery(&ttoInfo)
	if err, rettoInfo := service.GetTtoInfo(ttoInfo.ID); err != nil {
		invoker.Logger.Error("查询失败!", zap.Any("err", err))
		resp.FailWithMessage("查询失败", c)
	} else {
		resp.OkWithData(gin.H{"rettoInfo": rettoInfo}, c)
	}
}

// @Tags TtoInfo
// @Summary 分页获取TtoInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TtoInfoSearch true "分页获取TtoInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ttoInfo/getTtoInfoList [get]
func GetTtoInfoList(c *gin.Context) {
	var pageInfo request.TtoInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTtoInfoInfoList(pageInfo); err != nil {
		invoker.Logger.Error("获取失败", zap.Any("err", err))
		resp.FailWithMessage("获取失败", c)
	} else {
		resp.OkWithDetailed(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
