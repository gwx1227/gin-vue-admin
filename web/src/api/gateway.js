import service from '@/utils/request'

// @Tags Gateway
// @Summary 创建Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gateway true "创建Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gateway/createGateway [post]
export const createGateway = (data) => {
     return service({
         url: "/gateway/createGateway",
         method: 'post',
         data
     })
 }


// @Tags Gateway
// @Summary 删除Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gateway true "删除Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gateway/deleteGateway [delete]
 export const deleteGateway = (data) => {
     return service({
         url: "/gateway/deleteGateway",
         method: 'delete',
         data
     })
 }

// @Tags Gateway
// @Summary 删除Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gateway/deleteGateway [delete]
 export const deleteGatewayByIds = (data) => {
     return service({
         url: "/gateway/deleteGatewayByIds",
         method: 'delete',
         data
     })
 }

// @Tags Gateway
// @Summary 更新Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gateway true "更新Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gateway/updateGateway [put]
 export const updateGateway = (data) => {
     return service({
         url: "/gateway/updateGateway",
         method: 'put',
         data
     })
 }


// @Tags Gateway
// @Summary 用id查询Gateway
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Gateway true "用id查询Gateway"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gateway/findGateway [get]
 export const findGateway = (params) => {
     return service({
         url: "/gateway/findGateway",
         method: 'get',
         params
     })
 }


// @Tags Gateway
// @Summary 分页获取Gateway列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Gateway列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gateway/getGatewayList [get]
 export const getGatewayList = (params) => {
     return service({
         url: "/gateway/getGatewayList",
         method: 'get',
         params
     })
 }