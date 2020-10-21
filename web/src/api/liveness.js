import service from '@/utils/request'

// @Tags Liveness
// @Summary 创建Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Liveness true "创建Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liveness/createLiveness [post]
export const createLiveness = (data) => {
     return service({
         url: "/liveness/createLiveness",
         method: 'post',
         data
     })
 }


// @Tags Liveness
// @Summary 删除Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Liveness true "删除Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liveness/deleteLiveness [delete]
 export const deleteLiveness = (data) => {
     return service({
         url: "/liveness/deleteLiveness",
         method: 'delete',
         data
     })
 }

// @Tags Liveness
// @Summary 删除Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /liveness/deleteLiveness [delete]
 export const deleteLivenessByIds = (data) => {
     return service({
         url: "/liveness/deleteLivenessByIds",
         method: 'delete',
         data
     })
 }

// @Tags Liveness
// @Summary 更新Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Liveness true "更新Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /liveness/updateLiveness [put]
 export const updateLiveness = (data) => {
     return service({
         url: "/liveness/updateLiveness",
         method: 'put',
         data
     })
 }


// @Tags Liveness
// @Summary 用id查询Liveness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Liveness true "用id查询Liveness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /liveness/findLiveness [get]
 export const findLiveness = (params) => {
     return service({
         url: "/liveness/findLiveness",
         method: 'get',
         params
     })
 }


// @Tags Liveness
// @Summary 分页获取Liveness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Liveness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /liveness/getLivenessList [get]
 export const getLivenessList = (params) => {
     return service({
         url: "/liveness/getLivenessList",
         method: 'get',
         params
     })
 }