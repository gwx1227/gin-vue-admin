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

// @Tags Deploy
// @Summary 创建Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deploy true "创建Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/createDeploy [post]
func CreateDeploy(c *gin.Context) {
	var deploy model.Deploy
	_ = c.ShouldBindJSON(&deploy)
	err := service.CreateDeploy(deploy)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Deploy
// @Summary 删除Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deploy true "删除Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deploy/deleteDeploy [delete]
func DeleteDeploy(c *gin.Context) {
	var deploy model.Deploy
	_ = c.ShouldBindJSON(&deploy)
	err := service.DeleteDeploy(deploy)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Deploy
// @Summary 批量删除Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deploy/deleteDeployByIds [delete]
func DeleteDeployByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteDeployByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Deploy
// @Summary 更新Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deploy true "更新Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deploy/updateDeploy [put]
func UpdateDeploy(c *gin.Context) {
	var deploy model.Deploy
	_ = c.ShouldBindJSON(&deploy)
	err := service.UpdateDeploy(&deploy)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Deploy
// @Summary 用id查询Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deploy true "用id查询Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deploy/findDeploy [get]
func FindDeploy(c *gin.Context) {
	var deploy model.Deploy
	_ = c.ShouldBindQuery(&deploy)
	err, redeploy := service.GetDeploy(deploy.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redeploy": redeploy}, c)
	}
}

func FindDeployByAppId(c *gin.Context) {
	var deploy model.Deploy
	_ = c.ShouldBindQuery(&deploy)
	err, redeploy := service.GetDeployByAppId(deploy.AppId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redeploy": redeploy}, c)
	}
}

// @Tags Deploy
// @Summary 分页获取Deploy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeploySearch true "分页获取Deploy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/getDeployList [get]
func GetDeployList(c *gin.Context) {
	var pageInfo request.DeploySearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetDeployInfoList(pageInfo)
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
