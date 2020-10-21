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

// @Tags Alias
// @Summary 创建Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alias true "创建Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /alias/createAlias [post]
func CreateAlias(c *gin.Context) {
	var alias model.Alias
	_ = c.ShouldBindJSON(&alias)
	err := service.CreateAlias(alias)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Alias
// @Summary 删除Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alias true "删除Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alias/deleteAlias [delete]
func DeleteAlias(c *gin.Context) {
	var alias model.Alias
	_ = c.ShouldBindJSON(&alias)
	err := service.DeleteAlias(alias)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Alias
// @Summary 批量删除Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alias/deleteAliasByIds [delete]
func DeleteAliasByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAliasByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Alias
// @Summary 更新Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alias true "更新Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /alias/updateAlias [put]
func UpdateAlias(c *gin.Context) {
	var alias model.Alias
	_ = c.ShouldBindJSON(&alias)
	err := service.UpdateAlias(&alias)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Alias
// @Summary 用id查询Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alias true "用id查询Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /alias/findAlias [get]
func FindAlias(c *gin.Context) {
	var alias model.Alias
	_ = c.ShouldBindQuery(&alias)
	err, realias := service.GetAlias(alias.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"realias": realias}, c)
	}
}

// @Tags Alias
// @Summary 分页获取Alias列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AliasSearch true "分页获取Alias列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /alias/getAliasList [get]
func GetAliasList(c *gin.Context) {
	var pageInfo request.AliasSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAliasInfoList(pageInfo)
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
