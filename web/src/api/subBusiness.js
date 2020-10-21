import service from '@/utils/request'

// @Tags SubBusiness
// @Summary 创建SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubBusiness true "创建SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subBusiness/createSubBusiness [post]
export const createSubBusiness = (data) => {
     return service({
         url: "/subBusiness/createSubBusiness",
         method: 'post',
         data
     })
 }


// @Tags SubBusiness
// @Summary 删除SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubBusiness true "删除SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /subBusiness/deleteSubBusiness [delete]
 export const deleteSubBusiness = (data) => {
     return service({
         url: "/subBusiness/deleteSubBusiness",
         method: 'delete',
         data
     })
 }

// @Tags SubBusiness
// @Summary 删除SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /subBusiness/deleteSubBusiness [delete]
 export const deleteSubBusinessByIds = (data) => {
     return service({
         url: "/subBusiness/deleteSubBusinessByIds",
         method: 'delete',
         data
     })
 }

// @Tags SubBusiness
// @Summary 更新SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubBusiness true "更新SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /subBusiness/updateSubBusiness [put]
 export const updateSubBusiness = (data) => {
     return service({
         url: "/subBusiness/updateSubBusiness",
         method: 'put',
         data
     })
 }


// @Tags SubBusiness
// @Summary 用id查询SubBusiness
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubBusiness true "用id查询SubBusiness"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /subBusiness/findSubBusiness [get]
 export const findSubBusiness = (params) => {
     return service({
         url: "/subBusiness/findSubBusiness",
         method: 'get',
         params
     })
 }


// @Tags SubBusiness
// @Summary 分页获取SubBusiness列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SubBusiness列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subBusiness/getSubBusinessList [get]
 export const getSubBusinessList = (params) => {
     return service({
         url: "/subBusiness/getSubBusinessList",
         method: 'get',
         params
     })
 }