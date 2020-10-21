import service from '@/utils/request'

// @Tags Hpa
// @Summary 创建Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hpa true "创建Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hpa/createHpa [post]
export const createHpa = (data) => {
     return service({
         url: "/hpa/createHpa",
         method: 'post',
         data
     })
 }


// @Tags Hpa
// @Summary 删除Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hpa true "删除Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hpa/deleteHpa [delete]
 export const deleteHpa = (data) => {
     return service({
         url: "/hpa/deleteHpa",
         method: 'delete',
         data
     })
 }

// @Tags Hpa
// @Summary 删除Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hpa/deleteHpa [delete]
 export const deleteHpaByIds = (data) => {
     return service({
         url: "/hpa/deleteHpaByIds",
         method: 'delete',
         data
     })
 }

// @Tags Hpa
// @Summary 更新Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hpa true "更新Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hpa/updateHpa [put]
 export const updateHpa = (data) => {
     return service({
         url: "/hpa/updateHpa",
         method: 'put',
         data
     })
 }


// @Tags Hpa
// @Summary 用id查询Hpa
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Hpa true "用id查询Hpa"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hpa/findHpa [get]
 export const findHpa = (params) => {
     return service({
         url: "/hpa/findHpa",
         method: 'get',
         params
     })
 }


// @Tags Hpa
// @Summary 分页获取Hpa列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Hpa列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hpa/getHpaList [get]
 export const getHpaList = (params) => {
     return service({
         url: "/hpa/getHpaList",
         method: 'get',
         params
     })
 }