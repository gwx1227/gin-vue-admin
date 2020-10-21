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

// @Tags Business
// @Summary 创建Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Businesss true "创建Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /business/createBusiness [post]
func CreateBusiness(c *gin.Context) {
	var business model.Businesss
	_ = c.ShouldBindJSON(&business)
	err := service.CreateBusiness(business)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Business
// @Summary 删除Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Businesss true "删除Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /business/deleteBusiness [delete]
func DeleteBusiness(c *gin.Context) {
	var business model.Businesss
	_ = c.ShouldBindJSON(&business)
	err := service.DeleteBusiness(business)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Business
// @Summary 批量删除Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /business/deleteBusinessByIds [delete]
func DeleteBusinessByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteBusinessByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Business
// @Summary 更新Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Business true "更新Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /business/updateBusiness [put]
func UpdateBusiness(c *gin.Context) {
	var business model.Businesss
	_ = c.ShouldBindJSON(&business)
	err := service.UpdateBusiness(&business)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Business
// @Summary 用id查询Business
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Businesss true "用id查询Business"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /business/findBusiness [get]
func FindBusiness(c *gin.Context) {
	var business model.Businesss
	_ = c.ShouldBindQuery(&business)
	err, rebusiness := service.GetBusiness(business.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rebusiness": rebusiness}, c)
	}
}

// @Tags Business
// @Summary 分页获取Business列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BusinessSearch true "分页获取Business列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /business/getBusinessList [get]
func GetBusinessList(c *gin.Context) {
	var pageInfo request.BusinessSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetBusinessInfoList(pageInfo)
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
