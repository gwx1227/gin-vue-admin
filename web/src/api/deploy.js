import service from '@/utils/request'

// @Tags Deploy
// @Summary 创建Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deploy true "创建Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/createDeploy [post]
export const createDeploy = (data) => {
     return service({
         url: "/deploy/createDeploy",
         method: 'post',
         data
     })
 }


// @Tags Deploy
// @Summary 删除Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deploy true "删除Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deploy/deleteDeploy [delete]
 export const deleteDeploy = (data) => {
     return service({
         url: "/deploy/deleteDeploy",
         method: 'delete',
         data
     })
 }

// @Tags Deploy
// @Summary 删除Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deploy/deleteDeploy [delete]
 export const deleteDeployByIds = (data) => {
     return service({
         url: "/deploy/deleteDeployByIds",
         method: 'delete',
         data
     })
 }

// @Tags Deploy
// @Summary 更新Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deploy true "更新Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deploy/updateDeploy [put]
 export const updateDeploy = (data) => {
     return service({
         url: "/deploy/updateDeploy",
         method: 'put',
         data
     })
 }


// @Tags Deploy
// @Summary 用id查询Deploy
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Deploy true "用id查询Deploy"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deploy/findDeploy [get]
 export const findDeploy = (params) => {
     return service({
         url: "/deploy/findDeploy",
         method: 'get',
         params
     })
 }
export const findDeployByAppId = (params) => {
     return service({
         url: "/deploy/findDeployByAppId",
         method: 'get',
         params
     })
 }


// @Tags Deploy
// @Summary 分页获取Deploy列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Deploy列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/getDeployList [get]
 export const getDeployList = (params) => {
     return service({
         url: "/deploy/getDeployList",
         method: 'get',
         params
     })
 }