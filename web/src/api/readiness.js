import service from '@/utils/request'

// @Tags Readiness
// @Summary 创建Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Readiness true "创建Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /readiness/createReadiness [post]
export const createReadiness = (data) => {
     return service({
         url: "/readiness/createReadiness",
         method: 'post',
         data
     })
 }


// @Tags Readiness
// @Summary 删除Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Readiness true "删除Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /readiness/deleteReadiness [delete]
 export const deleteReadiness = (data) => {
     return service({
         url: "/readiness/deleteReadiness",
         method: 'delete',
         data
     })
 }

// @Tags Readiness
// @Summary 删除Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /readiness/deleteReadiness [delete]
 export const deleteReadinessByIds = (data) => {
     return service({
         url: "/readiness/deleteReadinessByIds",
         method: 'delete',
         data
     })
 }

// @Tags Readiness
// @Summary 更新Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Readiness true "更新Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /readiness/updateReadiness [put]
 export const updateReadiness = (data) => {
     return service({
         url: "/readiness/updateReadiness",
         method: 'put',
         data
     })
 }


// @Tags Readiness
// @Summary 用id查询Readiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Readiness true "用id查询Readiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /readiness/findReadiness [get]
 export const findReadiness = (params) => {
     return service({
         url: "/readiness/findReadiness",
         method: 'get',
         params
     })
 }


// @Tags Readiness
// @Summary 分页获取Readiness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Readiness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /readiness/getReadinessList [get]
 export const getReadinessList = (params) => {
     return service({
         url: "/readiness/getReadinessList",
         method: 'get',
         params
     })
 }