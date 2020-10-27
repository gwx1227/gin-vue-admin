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

// @Tags Namespaces
// @Summary 创建Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespaces true "创建Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /namespaces/createNamespaces [post]
func CreateNamespaces(c *gin.Context) {
	var namespaces model.Namespaces
	_ = c.ShouldBindJSON(&namespaces)
	err := service.CreateNamespaces(namespaces)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Namespaces
// @Summary 删除Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespaces true "删除Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /namespaces/deleteNamespaces [delete]
func DeleteNamespaces(c *gin.Context) {
	var namespaces model.Namespaces
	_ = c.ShouldBindJSON(&namespaces)
	err := service.DeleteNamespaces(namespaces)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Namespaces
// @Summary 批量删除Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /namespaces/deleteNamespacesByIds [delete]
func DeleteNamespacesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteNamespacesByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Namespaces
// @Summary 更新Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespaces true "更新Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /namespaces/updateNamespaces [put]
func UpdateNamespaces(c *gin.Context) {
	var namespaces model.Namespaces
	_ = c.ShouldBindJSON(&namespaces)
	err := service.UpdateNamespaces(&namespaces)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Namespaces
// @Summary 用id查询Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespaces true "用id查询Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /namespaces/findNamespaces [get]
func FindNamespaces(c *gin.Context) {
	var namespaces model.Namespaces
	_ = c.ShouldBindQuery(&namespaces)
	err, renamespaces := service.GetNamespaces(namespaces.ID)
	if err != nil {
		fmt.Println()
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"renamespaces": renamespaces}, c)
	}
}

// @Tags Namespaces
// @Summary 分页获取Namespaces列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NamespacesSearch true "分页获取Namespaces列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /namespaces/getNamespacesList [get]
func GetNamespacesList(c *gin.Context) {
	var pageInfo request.NamespacesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetNamespacesInfoList(pageInfo)
	userId := c.Request.Header.Get("X-User-Id")
	fmt.Printf("当前用户ID: %v\n", userId)
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
