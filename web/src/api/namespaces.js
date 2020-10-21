import service from '@/utils/request'

// @Tags Namespaces
// @Summary 创建Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespaces true "创建Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /namespaces/createNamespaces [post]
export const createNamespaces = (data) => {
     return service({
         url: "/namespaces/createNamespaces",
         method: 'post',
         data
     })
 }


// @Tags Namespaces
// @Summary 删除Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespaces true "删除Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /namespaces/deleteNamespaces [delete]
 export const deleteNamespaces = (data) => {
     return service({
         url: "/namespaces/deleteNamespaces",
         method: 'delete',
         data
     })
 }

// @Tags Namespaces
// @Summary 删除Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /namespaces/deleteNamespaces [delete]
 export const deleteNamespacesByIds = (data) => {
     return service({
         url: "/namespaces/deleteNamespacesByIds",
         method: 'delete',
         data
     })
 }

// @Tags Namespaces
// @Summary 更新Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespaces true "更新Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /namespaces/updateNamespaces [put]
 export const updateNamespaces = (data) => {
     return service({
         url: "/namespaces/updateNamespaces",
         method: 'put',
         data
     })
 }


// @Tags Namespaces
// @Summary 用id查询Namespaces
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespaces true "用id查询Namespaces"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /namespaces/findNamespaces [get]
 export const findNamespaces = (params) => {
     return service({
         url: "/namespaces/findNamespaces",
         method: 'get',
         params
     })
 }


// @Tags Namespaces
// @Summary 分页获取Namespaces列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Namespaces列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /namespaces/getNamespacesList [get]
 export const getNamespacesList = (params) => {
     return service({
         url: "/namespaces/getNamespacesList",
         method: 'get',
         params
     })
 }