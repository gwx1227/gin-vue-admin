package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
)

// @title    CreateGateway
// @description   create a Gateway
// @param     gateway               model.Gateway
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateGateway(gateway model.Gateway) (err error) {
	if gateway.WeightCanary + gateway.WeightOnline != 100 {
		err = errors.New("权重和应为100,请核实!")
	}
	err = global.GVA_DB.Create(&gateway).Error
	return err
}

// @title    DeleteGateway
// @description   delete a Gateway
// @auth                     （2020/04/05  20:22）
// @param     gateway               model.Gateway
// @return                    error

func DeleteGateway(gateway model.Gateway) (err error) {
	err = global.GVA_DB.Delete(gateway).Error
	return err
}

// @title    DeleteGatewayByIds
// @description   delete Gateways
// @auth                     （2020/04/05  20:22）
// @param     gateway               model.Gateway
// @return                    error

func DeleteGatewayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Gateway{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateGateway
// @description   update a Gateway
// @param     gateway          *model.Gateway
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateGateway(gateway *model.Gateway) (err error) {
	if gateway.WeightCanary + gateway.WeightOnline != 100 {
		err = errors.New("权重和应为100,请核实!")
	}
	err = global.GVA_DB.Where("app_id = ? and id = ?", gateway.AppId, gateway.ID).Updates(gateway).Error
	return err
}

// @title    GetGateway
// @description   get the info of a Gateway
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Gateway        Gateway

func GetGateway(id uint) (err error, gateway model.Gateway) {
	err = global.GVA_DB.Where("id = ?", id).First(&gateway).Error
	return
}

func GetGatewayByAppId(appId uint) (err error, gatewayData []resp.GatewayData) {
	err = global.GVA_DB.Table("gateway").Select("hosts,paths,weight_online,weight_canary").Where("app_id = ? and deleted_at is NULL", appId).Find(&gatewayData).Error
	return
}
func GetGatewayHostByAppId(appId uint) (err error, gatewayData []resp.GatewayHosts) {
	err = global.GVA_DB.Table("gateway").Select("hosts").Group("hosts").Where("app_id = ? and deleted_at is NULL", appId).Find(&gatewayData).Error
	return
}

// @title    GetGatewayInfoList
// @description   get Gateway list by pagination, 分页获取Gateway
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetGatewayInfoList(info request.GatewaySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Gateway{})
	var gateways []model.Gateway
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("app_id = ?", info.AppId).Count(&total).Error
	err = db.Where("app_id = ?", info.AppId).Limit(limit).Offset(offset).Find(&gateways).Error
	return err, gateways, total
}
