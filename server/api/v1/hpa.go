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

// @Tags Hpa
// @Summary 创建Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hpa true "创建Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hpa/createHpa [post]
func CreateHpa(c *gin.Context) {
	var hpa model.Hpa
	_ = c.ShouldBindJSON(&hpa)
	err := service.CreateHpa(hpa)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Hpa
// @Summary 删除Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hpa true "删除Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hpa/deleteHpa [delete]
func DeleteHpa(c *gin.Context) {
	var hpa model.Hpa
	_ = c.ShouldBindJSON(&hpa)
	err := service.DeleteHpa(hpa)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Hpa
// @Summary 批量删除Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hpa/deleteHpaByIds [delete]
func DeleteHpaByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteHpaByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Hpa
// @Summary 更新Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hpa true "更新Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hpa/updateHpa [put]
func UpdateHpa(c *gin.Context) {
	var hpa model.Hpa
	_ = c.ShouldBindJSON(&hpa)
	err := service.UpdateHpa(&hpa)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Hpa
// @Summary 用id查询Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hpa true "用id查询Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hpa/findHpa [get]
func FindHpa(c *gin.Context) {
	var hpa model.Hpa
	_ = c.ShouldBindQuery(&hpa)
	err, rehpa := service.GetHpa(hpa.AppId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rehpa": rehpa}, c)
	}
}

// @Tags Hpa
// @Summary 分页获取Hpa列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HpaSearch true "分页获取Hpa列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hpa/getHpaList [get]
func GetHpaList(c *gin.Context) {
	var pageInfo request.HpaSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetHpaInfoList(pageInfo)
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
