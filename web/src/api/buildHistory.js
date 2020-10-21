import service from '@/utils/request'

// @Tags BuildHistory
// @Summary 创建BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildHistory true "创建BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildHistory/createBuildHistory [post]
export const createBuildHistory = (data) => {
     return service({
         url: "/buildHistory/createBuildHistory",
         method: 'post',
         data
     })
 }


// @Tags BuildHistory
// @Summary 删除BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildHistory true "删除BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildHistory/deleteBuildHistory [delete]
 export const deleteBuildHistory = (data) => {
     return service({
         url: "/buildHistory/deleteBuildHistory",
         method: 'delete',
         data
     })
 }

// @Tags BuildHistory
// @Summary 删除BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /buildHistory/deleteBuildHistory [delete]
 export const deleteBuildHistoryByIds = (data) => {
     return service({
         url: "/buildHistory/deleteBuildHistoryByIds",
         method: 'delete',
         data
     })
 }

// @Tags BuildHistory
// @Summary 更新BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildHistory true "更新BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /buildHistory/updateBuildHistory [put]
 export const updateBuildHistory = (data) => {
     return service({
         url: "/buildHistory/updateBuildHistory",
         method: 'put',
         data
     })
 }


// @Tags BuildHistory
// @Summary 用id查询BuildHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BuildHistory true "用id查询BuildHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /buildHistory/findBuildHistory [get]
 export const findBuildHistory = (params) => {
     return service({
         url: "/buildHistory/findBuildHistory",
         method: 'get',
         params
     })
 }


// @Tags BuildHistory
// @Summary 分页获取BuildHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取BuildHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /buildHistory/getBuildHistoryList [get]
 export const getBuildHistoryList = (params) => {
     return service({
         url: "/buildHistory/getBuildHistoryList",
         method: 'get',
         params
     })
 }