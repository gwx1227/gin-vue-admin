import service from '@/utils/request'

// @Tags DeployHistory
// @Summary 创建DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "创建DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deployHistory/createDeployHistory [post]
export const createDeployHistory = (data) => {
     return service({
         url: "/deployHistory/createDeployHistory",
         method: 'post',
         data
     })
 }


// @Tags DeployHistory
// @Summary 删除DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "删除DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deployHistory/deleteDeployHistory [delete]
 export const deleteDeployHistory = (data) => {
     return service({
         url: "/deployHistory/deleteDeployHistory",
         method: 'delete',
         data
     })
 }

// @Tags DeployHistory
// @Summary 删除DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deployHistory/deleteDeployHistory [delete]
 export const deleteDeployHistoryByIds = (data) => {
     return service({
         url: "/deployHistory/deleteDeployHistoryByIds",
         method: 'delete',
         data
     })
 }

// @Tags DeployHistory
// @Summary 更新DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "更新DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deployHistory/updateDeployHistory [put]
 export const updateDeployHistory = (data) => {
     return service({
         url: "/deployHistory/updateDeployHistory",
         method: 'put',
         data
     })
 }


// @Tags DeployHistory
// @Summary 用id查询DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "用id查询DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deployHistory/findDeployHistory [get]
 export const findDeployHistory = (params) => {
     return service({
         url: "/deployHistory/findDeployHistory",
         method: 'get',
         params
     })
 }


// @Tags DeployHistory
// @Summary 分页获取DeployHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DeployHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deployHistory/getDeployHistoryList [get]
 export const getDeployHistoryList = (params) => {
     return service({
         url: "/deployHistory/getDeployHistoryList",
         method: 'get',
         params
     })
 }