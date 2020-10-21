package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Gateway
// @Summary 创建Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gateway true "创建Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gateway/createGateway [post]
func CreateGateway(c *gin.Context) {
	var gateway model.Gateway
	_ = c.ShouldBindJSON(&gateway)
	err := service.CreateGateway(gateway)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Gateway
// @Summary 删除Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gateway true "删除Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gateway/deleteGateway [delete]
func DeleteGateway(c *gin.Context) {
	var gateway model.Gateway
	_ = c.ShouldBindJSON(&gateway)
	err := service.DeleteGateway(gateway)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Gateway
// @Summary 批量删除Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gateway/deleteGatewayByIds [delete]
func DeleteGatewayByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteGatewayByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Gateway
// @Summary 更新Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gateway true "更新Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gateway/updateGateway [put]
func UpdateGateway(c *gin.Context) {
	var gateway model.Gateway
	_ = c.ShouldBindJSON(&gateway)
	err := service.UpdateGateway(&gateway)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Gateway
// @Summary 用id查询Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gateway true "用id查询Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gateway/findGateway [get]
func FindGateway(c *gin.Context) {
	var gateway model.Gateway
	_ = c.ShouldBindQuery(&gateway)
	err, regateway := service.GetGateway(gateway.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"regateway": regateway}, c)
	}
}

// @Tags Gateway
// @Summary 分页获取Gateway列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GatewaySearch true "分页获取Gateway列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gateway/getGatewayList [get]
func GetGatewayList(c *gin.Context) {
	var pageInfo request.GatewaySearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetGatewayInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
