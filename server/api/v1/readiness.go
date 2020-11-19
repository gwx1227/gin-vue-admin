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

// @Tags Readiness
// @Summary 创建Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Readiness true "创建Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /readiness/createReadiness [post]
func CreateReadiness(c *gin.Context) {
	var readiness model.Readiness
	_ = c.ShouldBindJSON(&readiness)
	err := service.CreateReadiness(readiness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Readiness
// @Summary 删除Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Readiness true "删除Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /readiness/deleteReadiness [delete]
func DeleteReadiness(c *gin.Context) {
	var readiness model.Readiness
	_ = c.ShouldBindJSON(&readiness)
	err := service.DeleteReadiness(readiness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Readiness
// @Summary 批量删除Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /readiness/deleteReadinessByIds [delete]
func DeleteReadinessByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteReadinessByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Readiness
// @Summary 更新Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Readiness true "更新Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /readiness/updateReadiness [put]
func UpdateReadiness(c *gin.Context) {
	var readiness model.Readiness
	_ = c.ShouldBindJSON(&readiness)
	err := service.UpdateReadiness(&readiness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Readiness
// @Summary 用id查询Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Readiness true "用id查询Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /readiness/findReadiness [get]
func FindReadiness(c *gin.Context) {
	var readiness model.Readiness
	_ = c.ShouldBindQuery(&readiness)
	err, rereadiness := service.GetReadiness(readiness.AppId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rereadiness": rereadiness}, c)
	}
}

// @Tags Readiness
// @Summary 分页获取Readiness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ReadinessSearch true "分页获取Readiness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /readiness/getReadinessList [get]
func GetReadinessList(c *gin.Context) {
	var pageInfo request.ReadinessSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetReadinessInfoList(pageInfo)
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
