import service from '@/utils/request'

// @Tags Apps
// @Summary 创建Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apps true "创建Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apps/createApps [post]
export const createApps = (data) => {
     return service({
         url: "/apps/createApps",
         method: 'post',
         data
     })
 }


// @Tags Apps
// @Summary 删除Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apps true "删除Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apps/deleteApps [delete]
 export const deleteApps = (data) => {
     return service({
         url: "/apps/deleteApps",
         method: 'delete',
         data
     })
 }

// @Tags Apps
// @Summary 删除Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apps/deleteApps [delete]
 export const deleteAppsByIds = (data) => {
     return service({
         url: "/apps/deleteAppsByIds",
         method: 'delete',
         data
     })
 }

// @Tags Apps
// @Summary 更新Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apps true "更新Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apps/updateApps [put]
 export const updateApps = (data) => {
     return service({
         url: "/apps/updateApps",
         method: 'put',
         data
     })
 }

export const updateAppsSwitch = (data) => {
     return service({
         url: "/apps/updateAppsSwitch",
         method: 'put',
         data
     })
} 

// @Tags Apps
// @Summary 用id查询Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apps true "用id查询Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apps/findApps [get]
 export const findApps = (params) => {
     return service({
         url: "/apps/findApps",
         method: 'get',
         params
     })
 }


// @Tags Apps
// @Summary 分页获取Apps列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Apps列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apps/getAppsList [get]
 export const getAppsList = (params) => {
     return service({
         url: "/apps/getAppsList",
         method: 'get',
         params
     })
 }

export const  getAppsListByNamespaceId = (params) => {
    return service({
        url: "/apps/getAppsListByNamespaceId",
        method: 'get',
        params
    })
}