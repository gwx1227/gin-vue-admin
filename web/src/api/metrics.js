import service from '@/utils/request'

// @Tags Metrics
// @Summary 创建Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Metrics true "创建Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /metrics/createMetrics [post]
export const createMetrics = (data) => {
     return service({
         url: "/metrics/createMetrics",
         method: 'post',
         data
     })
 }


// @Tags Metrics
// @Summary 删除Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Metrics true "删除Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /metrics/deleteMetrics [delete]
 export const deleteMetrics = (data) => {
     return service({
         url: "/metrics/deleteMetrics",
         method: 'delete',
         data
     })
 }

// @Tags Metrics
// @Summary 删除Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /metrics/deleteMetrics [delete]
 export const deleteMetricsByIds = (data) => {
     return service({
         url: "/metrics/deleteMetricsByIds",
         method: 'delete',
         data
     })
 }

// @Tags Metrics
// @Summary 更新Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Metrics true "更新Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /metrics/updateMetrics [put]
 export const updateMetrics = (data) => {
     return service({
         url: "/metrics/updateMetrics",
         method: 'put',
         data
     })
 }


// @Tags Metrics
// @Summary 用id查询Metrics
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Metrics true "用id查询Metrics"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /metrics/findMetrics [get]
 export const findMetrics = (params) => {
     return service({
         url: "/metrics/findMetrics",
         method: 'get',
         params
     })
 }


// @Tags Metrics
// @Summary 分页获取Metrics列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Metrics列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /metrics/getMetricsList [get]
 export const getMetricsList = (params) => {
     return service({
         url: "/metrics/getMetricsList",
         method: 'get',
         params
     })
 }