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

// @Tags SysConfig
// @Summary 创建SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "创建SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysConfig/createSysConfig [post]
func CreateSysConfig(c *gin.Context) {
	var sysConfig model.SysConfig
	_ = c.ShouldBindJSON(&sysConfig)
	err := service.CreateSysConfig(sysConfig)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysConfig
// @Summary 删除SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "删除SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysConfig/deleteSysConfig [delete]
func DeleteSysConfig(c *gin.Context) {
	var sysConfig model.SysConfig
	_ = c.ShouldBindJSON(&sysConfig)
	err := service.DeleteSysConfig(sysConfig)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysConfig
// @Summary 批量删除SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysConfig/deleteSysConfigByIds [delete]
func DeleteSysConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSysConfigByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysConfig
// @Summary 更新SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "更新SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysConfig/updateSysConfig [put]
func UpdateSysConfig(c *gin.Context) {
	var sysConfig model.SysConfig
	_ = c.ShouldBindJSON(&sysConfig)
	err := service.UpdateSysConfig(&sysConfig)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysConfig
// @Summary 用id查询SysConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "用id查询SysConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysConfig/findSysConfig [get]
func FindSysConfig(c *gin.Context) {
	var sysConfig model.SysConfig
	_ = c.ShouldBindQuery(&sysConfig)
	err, resysConfig := service.GetSysConfig(sysConfig.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resysConfig": resysConfig}, c)
	}
}

// @Tags SysConfig
// @Summary 分页获取SysConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SysConfigSearch true "分页获取SysConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysConfig/getSysConfigList [get]
func GetSysConfigList(c *gin.Context) {
	var pageInfo request.SysConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSysConfigInfoList(pageInfo)
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
