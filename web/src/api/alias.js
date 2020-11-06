import service from '@/utils/request'

// @Tags Alias
// @Summary 创建Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alias true "创建Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /alias/createAlias [post]
export const createAlias = (data) => {
     return service({
         url: "/alias/createAlias",
         method: 'post',
         data
     })
 }


// @Tags Alias
// @Summary 删除Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alias true "删除Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alias/deleteAlias [delete]
 export const deleteAlias = (data) => {
     return service({
         url: "/alias/deleteAlias",
         method: 'delete',
         data
     })
 }

// @Tags Alias
// @Summary 删除Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alias/deleteAlias [delete]
 export const deleteAliasByIds = (data) => {
     return service({
         url: "/alias/deleteAliasByIds",
         method: 'delete',
         data
     })
 }

// @Tags Alias
// @Summary 更新Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alias true "更新Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /alias/updateAlias [put]
 export const updateAlias = (data) => {
     return service({
         url: "/alias/updateAlias",
         method: 'put',
         data
     })
 }


// @Tags Alias
// @Summary 用id查询Alias
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Alias true "用id查询Alias"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /alias/findAlias [get]
 export const findAlias = (params) => {
     return service({
         url: "/alias/findAlias",
         method: 'get',
         params
     })
 }


// @Tags Alias
// @Summary 分页获取Alias列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Alias列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /alias/getAliasList [get]
 export const getAliasList = (params) => {
     return service({
         url: "/alias/getAliasList",
         method: 'get',
         params
     })
}

 export const getAliasListByAppId = (params) => {
     return service({
         url: "/alias/getAliasListByAppId",
         method: 'get',
         params
     })
 }