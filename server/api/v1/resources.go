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

// @Tags Resources
// @Summary 创建Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resources true "创建Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resources/createResources [post]
func CreateResources(c *gin.Context) {
	var resources model.Resources
	_ = c.ShouldBindJSON(&resources)
	err := service.CreateResources(resources)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Resources
// @Summary 删除Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resources true "删除Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resources/deleteResources [delete]
func DeleteResources(c *gin.Context) {
	var resources model.Resources
	_ = c.ShouldBindJSON(&resources)
	err := service.DeleteResources(resources)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Resources
// @Summary 批量删除Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resources/deleteResourcesByIds [delete]
func DeleteResourcesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteResourcesByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Resources
// @Summary 更新Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resources true "更新Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /resources/updateResources [put]
func UpdateResources(c *gin.Context) {
	var resources model.Resources
	_ = c.ShouldBindJSON(&resources)
	err := service.UpdateResources(&resources)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Resources
// @Summary 用id查询Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resources true "用id查询Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /resources/findResources [get]
func FindResources(c *gin.Context) {
	var resources model.Resources
	_ = c.ShouldBindQuery(&resources)
	err, reresources := service.GetResources(resources.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reresources": reresources}, c)
	}
}

// @Tags Resources
// @Summary 分页获取Resources列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ResourcesSearch true "分页获取Resources列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resources/getResourcesList [get]
func GetResourcesList(c *gin.Context) {
	var pageInfo request.ResourcesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetResourcesInfoList(pageInfo)
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
