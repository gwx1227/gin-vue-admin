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

// @Tags SubBusiness
// @Summary 创建SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubBusiness true "创建SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subBusiness/createSubBusiness [post]
func CreateSubBusiness(c *gin.Context) {
	var subBusiness model.SubBusiness
	_ = c.ShouldBindJSON(&subBusiness)
	err := service.CreateSubBusiness(subBusiness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SubBusiness
// @Summary 删除SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubBusiness true "删除SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /subBusiness/deleteSubBusiness [delete]
func DeleteSubBusiness(c *gin.Context) {
	var subBusiness model.SubBusiness
	_ = c.ShouldBindJSON(&subBusiness)
	err := service.DeleteSubBusiness(subBusiness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SubBusiness
// @Summary 批量删除SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /subBusiness/deleteSubBusinessByIds [delete]
func DeleteSubBusinessByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSubBusinessByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SubBusiness
// @Summary 更新SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubBusiness true "更新SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /subBusiness/updateSubBusiness [put]
func UpdateSubBusiness(c *gin.Context) {
	var subBusiness model.SubBusiness
	_ = c.ShouldBindJSON(&subBusiness)
	err := service.UpdateSubBusiness(&subBusiness)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SubBusiness
// @Summary 用id查询SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubBusiness true "用id查询SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /subBusiness/findSubBusiness [get]
func FindSubBusiness(c *gin.Context) {
	var subBusiness model.SubBusiness
	_ = c.ShouldBindQuery(&subBusiness)
	err, resubBusiness := service.GetSubBusiness(subBusiness.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resubBusiness": resubBusiness}, c)
	}
}

// @Tags SubBusiness
// @Summary 分页获取SubBusiness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SubBusinessSearch true "分页获取SubBusiness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subBusiness/getSubBusinessList [get]
func GetSubBusinessList(c *gin.Context) {
	var pageInfo request.SubBusinessSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSubBusinessInfoList(pageInfo)
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
