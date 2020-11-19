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

// @Tags Metrics
// @Summary 创建Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Metrics true "创建Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /metrics/createMetrics [post]
func CreateMetrics(c *gin.Context) {
	var metrics model.Metrics
	_ = c.ShouldBindJSON(&metrics)
	err := service.CreateMetrics(metrics)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Metrics
// @Summary 删除Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Metrics true "删除Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /metrics/deleteMetrics [delete]
func DeleteMetrics(c *gin.Context) {
	var metrics model.Metrics
	_ = c.ShouldBindJSON(&metrics)
	err := service.DeleteMetrics(metrics)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Metrics
// @Summary 批量删除Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /metrics/deleteMetricsByIds [delete]
func DeleteMetricsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteMetricsByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Metrics
// @Summary 更新Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Metrics true "更新Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /metrics/updateMetrics [put]
func UpdateMetrics(c *gin.Context) {
	var metrics model.Metrics
	_ = c.ShouldBindJSON(&metrics)
	fmt.Printf("....... %v\n\n", metrics.Port)
	err := service.UpdateMetrics(&metrics)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Metrics
// @Summary 用id查询Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Metrics true "用id查询Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /metrics/findMetrics [get]
func FindMetrics(c *gin.Context) {
	var metrics model.Metrics
	_ = c.ShouldBindQuery(&metrics)
	err, remetrics := service.GetMetrics(metrics.AppId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"remetrics": remetrics}, c)
	}
}

// @Tags Metrics
// @Summary 分页获取Metrics列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MetricsSearch true "分页获取Metrics列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /metrics/getMetricsList [get]
func GetMetricsList(c *gin.Context) {
	var pageInfo request.MetricsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetMetricsInfoList(pageInfo)
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
