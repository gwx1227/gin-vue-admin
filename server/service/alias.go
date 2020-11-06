package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateAlias
// @description   create a Alias
// @param     alias               model.Alias
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateAlias(alias model.Alias) (err error) {
	err = global.GVA_DB.Create(&alias).Error
	return err
}

// @title    DeleteAlias
// @description   delete a Alias
// @auth                     （2020/04/05  20:22）
// @param     alias               model.Alias
// @return                    error

func DeleteAlias(alias model.Alias) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", alias.ID).Delete(alias).Error
	return err
}

// @title    DeleteAliasByIds
// @description   delete Aliass
// @auth                     （2020/04/05  20:22）
// @param     alias               model.Alias
// @return                    error

func DeleteAliasByIds(ids request.IdsReq) (err error) {

	err = global.GVA_DB.Debug().Delete(&[]model.Alias{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateAlias
// @description   update a Alias
// @param     alias          *model.Alias
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateAlias(alias *model.Alias) (err error) {
	//err = global.GVA_DB.Debug().Save(alias).Error
	err = global.GVA_DB.Debug().Where("app_id = ? and hostname = ?", alias.AppId,alias.Hostname).Updates(model.Alias{Ip:alias.Ip}).Error
	return err
}

// @title    GetAlias
// @description   get the info of a Alias
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Alias        Alias

func GetAlias(id uint) (err error, alias model.Alias) {
	err = global.GVA_DB.Where("id = ?", id).First(&alias).Error
	return
}
func GetAliasByAppId(app_id uint) (err error, alias model.Alias) {
	err = global.GVA_DB.Where("app_id = ?", app_id).First(&alias).Error
	return
}

// @title    GetAliasInfoList
// @description   get Alias list by pagination, 分页获取Alias
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAliasInfoList(info request.AliasSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Alias{})
	var aliass []model.Alias
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&aliass).Error
	return err, aliass, total
}


func GetAliasInfoListByAppId(info request.AliasSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Alias{})
	var aliass []model.Alias
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Debug().Where("app_id = ?", info.AppId).Count(&total).Error
	err = db.Debug().Limit(limit).Offset(offset).Where("app_id = ?", info.AppId).Find(&aliass).Error
	return err, aliass, total
}
