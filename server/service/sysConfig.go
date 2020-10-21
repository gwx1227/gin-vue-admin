package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSysConfig
// @description   create a SysConfig
// @param     sysConfig               model.SysConfig
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSysConfig(sysConfig model.SysConfig) (err error) {
	err = global.GVA_DB.Create(&sysConfig).Error
	return err
}

// @title    DeleteSysConfig
// @description   delete a SysConfig
// @auth                     （2020/04/05  20:22）
// @param     sysConfig               model.SysConfig
// @return                    error

func DeleteSysConfig(sysConfig model.SysConfig) (err error) {
	err = global.GVA_DB.Delete(sysConfig).Error
	return err
}

// @title    DeleteSysConfigByIds
// @description   delete SysConfigs
// @auth                     （2020/04/05  20:22）
// @param     sysConfig               model.SysConfig
// @return                    error

func DeleteSysConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SysConfig{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateSysConfig
// @description   update a SysConfig
// @param     sysConfig          *model.SysConfig
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSysConfig(sysConfig *model.SysConfig) (err error) {
	err = global.GVA_DB.Save(sysConfig).Error
	return err
}

// @title    GetSysConfig
// @description   get the info of a SysConfig
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SysConfig        SysConfig

func GetSysConfig(id uint) (err error, sysConfig model.SysConfig) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysConfig).Error
	return
}

// @title    GetSysConfigInfoList
// @description   get SysConfig list by pagination, 分页获取SysConfig
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSysConfigInfoList(info request.SysConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysConfig{})
	var sysConfigs []model.SysConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysConfigs).Error
	return err, sysConfigs, total
}
