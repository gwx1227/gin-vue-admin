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

// @Tags Config
// @Summary 创建Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Config true "创建Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /config/createConfig [post]
func CreateConfig(c *gin.Context) {
	var config model.Config
	_ = c.ShouldBindJSON(&config)
	err := service.CreateConfig(config)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Config
// @Summary 删除Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Config true "删除Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /config/deleteConfig [delete]
func DeleteConfig(c *gin.Context) {
	var config model.Config
	_ = c.ShouldBindJSON(&config)
	err := service.DeleteConfig(config)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Config
// @Summary 批量删除Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /config/deleteConfigByIds [delete]
func DeleteConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteConfigByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Config
// @Summary 更新Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Config true "更新Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /config/updateConfig [put]
func UpdateConfig(c *gin.Context) {
	var config model.Config
	_ = c.ShouldBindJSON(&config)
	err := service.UpdateConfig(&config)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Config
// @Summary 用id查询Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Config true "用id查询Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /config/findConfig [get]
func FindConfig(c *gin.Context) {
	var config model.Config
	_ = c.ShouldBindQuery(&config)
	err, reconfig := service.GetConfig(config.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reconfig": reconfig}, c)
	}
}

// @Tags Config
// @Summary 分页获取Config列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ConfigSearch true "分页获取Config列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /config/getConfigList [get]
func GetConfigList(c *gin.Context) {
	var pageInfo request.ConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetConfigInfoList(pageInfo)
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


func GetConfigListByAppId(c *gin.Context) {
	var pageInfo request.ConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetConfigInfoListByAppId(pageInfo)
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
