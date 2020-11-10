package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateConfig
// @description   create a Config
// @param     config               model.Config
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateConfig(config model.Config) (err error) {
	err = global.GVA_DB.Create(&config).Error
	return err
}

// @title    DeleteConfig
// @description   delete a Config
// @auth                     （2020/04/05  20:22）
// @param     config               model.Config
// @return                    error

func DeleteConfig(config model.Config) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", config.ID).Delete(config).Error
	return err
}

// @title    DeleteConfigByIds
// @description   delete Configs
// @auth                     （2020/04/05  20:22）
// @param     config               model.Config
// @return                    error

func DeleteConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Config{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateConfig
// @description   update a Config
// @param     config          *model.Config
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateConfig(config *model.Config) (err error) {
	err = global.GVA_DB.Debug().Where("app_id = ? and id = ?",config.AppId,config.ID).Updates(model.Config{Value: config.Value}).Error
	return err
}

// @title    GetConfig
// @description   get the info of a Config
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Config        Config

func GetConfig(id uint) (err error, config model.Config) {
	err = global.GVA_DB.Where("id = ?", id).First(&config).Error
	return
}

// @title    GetConfigInfoList
// @description   get Config list by pagination, 分页获取Config
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetConfigInfoList(info request.ConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Config{})
	var configs []model.Config
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&configs).Error
	return err, configs, total
}


func GetConfigInfoListByAppId(info request.ConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Config{})
	var configs []model.Config
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Debug().Where("app_id = ?", info.AppId).Count(&total).Error
	err = db.Debug().Limit(limit).Offset(offset).Where("app_id = ?", info.AppId).Find(&configs).Error
	return err, configs, total
}

