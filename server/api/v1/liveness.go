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

// @Tags Liveness
// @Summary 创建Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Liveness true "创建Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liveness/createLiveness [post]
func CreateLiveness(c *gin.Context) {
	var liveness model.Liveness
	_ = c.ShouldBindJSON(&liveness)
	err := service.CreateLiveness(liveness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Liveness
// @Summary 删除Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Liveness true "删除Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liveness/deleteLiveness [delete]
func DeleteLiveness(c *gin.Context) {
	var liveness model.Liveness
	_ = c.ShouldBindJSON(&liveness)
	err := service.DeleteLiveness(liveness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Liveness
// @Summary 批量删除Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liveness/deleteLivenessByIds [delete]
func DeleteLivenessByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLivenessByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Liveness
// @Summary 更新Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Liveness true "更新Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liveness/updateLiveness [put]
func UpdateLiveness(c *gin.Context) {
	var liveness model.Liveness
	_ = c.ShouldBindJSON(&liveness)
	err := service.UpdateLiveness(&liveness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Liveness
// @Summary 用id查询Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Liveness true "用id查询Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liveness/findLiveness [get]
func FindLiveness(c *gin.Context) {
	var liveness model.Liveness
	_ = c.ShouldBindQuery(&liveness)
	err, reliveness := service.GetLiveness(liveness.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reliveness": reliveness}, c)
	}
}

// @Tags Liveness
// @Summary 分页获取Liveness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LivenessSearch true "分页获取Liveness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liveness/getLivenessList [get]
func GetLivenessList(c *gin.Context) {
	var pageInfo request.LivenessSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLivenessInfoList(pageInfo)
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
