package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
)

// @title    CreateApps
// @description   create a Apps
// @param     apps               model.Apps
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateApps(apps model.Apps) (err error) {
	err = global.GVA_DB.Create(&apps).Error
	return err
}

// @title    DeleteApps
// @description   delete a Apps
// @auth                     （2020/04/05  20:22）
// @param     apps               model.Apps
// @return                    error

func DeleteApps(apps model.Apps) (err error) {
	err = global.GVA_DB.Delete(apps).Error
	return err
}

// @title    DeleteAppsByIds
// @description   delete Appss
// @auth                     （2020/04/05  20:22）
// @param     apps               model.Apps
// @return                    error

func DeleteAppsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Apps{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateApps
// @description   update a Apps
// @param     apps          *model.Apps
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateApps(apps *model.Apps) (err error) {
	err = global.GVA_DB.Save(apps).Error
	return err
}

// @title    GetApps
// @description   get the info of a Apps
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Apps        Apps

func GetApps(id uint) (err error, apps []response.AppInfo) {
	err = global.GVA_DB.Debug().Table("apps").Select("apps.*,business.business,sub_business.subbusiness ").Where("apps.id = ?", id).Joins("left join sub_business on sub_business.id = apps.subbusiness_id").Joins("left join business on business.id = apps.business_id").Scan(&apps).Error
	return
}

// @title    GetAppsInfoList
// @description   get Apps list by pagination, 分页获取Apps
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetAppsInfoList(info request.AppsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Apps{})
	var appss []model.Apps
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&appss).Error
	return err, appss, total
}

func GetAppsInfoListByNamespaceId(info request.AppsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Apps{})
	var appss []model.Apps
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Debug().Where("namespace_id = ? ", info.NamespaceId).Count(&total).Error
	if info.ID != 0 {
		err = db.Debug().Limit(limit).Offset(offset).Where("namespace_id = ? AND id = ?", info.NamespaceId, info.ID).Find(&appss).Error
	}else {
		err = db.Debug().Limit(limit).Offset(offset).Where("namespace_id = ?", info.NamespaceId).Find(&appss).Error
	}
	return err, appss, total
}

