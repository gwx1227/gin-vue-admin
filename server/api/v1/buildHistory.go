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

// @Tags BuildHistory
// @Summary 创建BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildHistory true "创建BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildHistory/createBuildHistory [post]
func CreateBuildHistory(c *gin.Context) {
	var buildHistory model.BuildHistory
	_ = c.ShouldBindJSON(&buildHistory)
	err := service.CreateBuildHistory(buildHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags BuildHistory
// @Summary 删除BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildHistory true "删除BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildHistory/deleteBuildHistory [delete]
func DeleteBuildHistory(c *gin.Context) {
	var buildHistory model.BuildHistory
	_ = c.ShouldBindJSON(&buildHistory)
	err := service.DeleteBuildHistory(buildHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BuildHistory
// @Summary 批量删除BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildHistory/deleteBuildHistoryByIds [delete]
func DeleteBuildHistoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteBuildHistoryByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BuildHistory
// @Summary 更新BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildHistory true "更新BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /buildHistory/updateBuildHistory [put]
func UpdateBuildHistory(c *gin.Context) {
	var buildHistory model.BuildHistory
	_ = c.ShouldBindJSON(&buildHistory)
	err := service.UpdateBuildHistory(&buildHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BuildHistory
// @Summary 用id查询BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildHistory true "用id查询BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /buildHistory/findBuildHistory [get]
func FindBuildHistory(c *gin.Context) {
	var buildHistory model.BuildHistory
	_ = c.ShouldBindQuery(&buildHistory)
	err, rebuildHistory := service.GetBuildHistory(buildHistory.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rebuildHistory": rebuildHistory}, c)
	}
}

// @Tags BuildHistory
// @Summary 分页获取BuildHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BuildHistorySearch true "分页获取BuildHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildHistory/getBuildHistoryList [get]
func GetBuildHistoryList(c *gin.Context) {
	var pageInfo request.BuildHistorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetBuildHistoryInfoList(pageInfo)
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
