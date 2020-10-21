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

// @Tags DeployHistory
// @Summary 创建DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "创建DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deployHistory/createDeployHistory [post]
func CreateDeployHistory(c *gin.Context) {
	var deployHistory model.DeployHistory
	_ = c.ShouldBindJSON(&deployHistory)
	err := service.CreateDeployHistory(deployHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DeployHistory
// @Summary 删除DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "删除DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deployHistory/deleteDeployHistory [delete]
func DeleteDeployHistory(c *gin.Context) {
	var deployHistory model.DeployHistory
	_ = c.ShouldBindJSON(&deployHistory)
	err := service.DeleteDeployHistory(deployHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DeployHistory
// @Summary 批量删除DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deployHistory/deleteDeployHistoryByIds [delete]
func DeleteDeployHistoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteDeployHistoryByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DeployHistory
// @Summary 更新DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "更新DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deployHistory/updateDeployHistory [put]
func UpdateDeployHistory(c *gin.Context) {
	var deployHistory model.DeployHistory
	_ = c.ShouldBindJSON(&deployHistory)
	err := service.UpdateDeployHistory(&deployHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DeployHistory
// @Summary 用id查询DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "用id查询DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deployHistory/findDeployHistory [get]
func FindDeployHistory(c *gin.Context) {
	var deployHistory model.DeployHistory
	_ = c.ShouldBindQuery(&deployHistory)
	err, redeployHistory := service.GetDeployHistory(deployHistory.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redeployHistory": redeployHistory}, c)
	}
}

// @Tags DeployHistory
// @Summary 分页获取DeployHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeployHistorySearch true "分页获取DeployHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deployHistory/getDeployHistoryList [get]
func GetDeployHistoryList(c *gin.Context) {
	var pageInfo request.DeployHistorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetDeployHistoryInfoList(pageInfo)
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
