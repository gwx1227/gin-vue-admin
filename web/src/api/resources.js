import service from '@/utils/request'

// @Tags Resources
// @Summary 创建Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resources true "创建Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resources/createResources [post]
export const createResources = (data) => {
     return service({
         url: "/resources/createResources",
         method: 'post',
         data
     })
 }


// @Tags Resources
// @Summary 删除Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resources true "删除Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resources/deleteResources [delete]
 export const deleteResources = (data) => {
     return service({
         url: "/resources/deleteResources",
         method: 'delete',
         data
     })
 }

// @Tags Resources
// @Summary 删除Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resources/deleteResources [delete]
 export const deleteResourcesByIds = (data) => {
     return service({
         url: "/resources/deleteResourcesByIds",
         method: 'delete',
         data
     })
 }

// @Tags Resources
// @Summary 更新Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resources true "更新Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /resources/updateResources [put]
 export const updateResources = (data) => {
     return service({
         url: "/resources/updateResources",
         method: 'put',
         data
     })
 }


// @Tags Resources
// @Summary 用id查询Resources
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Resources true "用id查询Resources"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /resources/findResources [get]
 export const findResources = (params) => {
     return service({
         url: "/resources/findResources",
         method: 'get',
         params
     })
 }


// @Tags Resources
// @Summary 分页获取Resources列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Resources列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resources/getResourcesList [get]
 export const getResourcesList = (params) => {
     return service({
         url: "/resources/getResourcesList",
         method: 'get',
         params
     })
 }