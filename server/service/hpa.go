package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateHpa
// @description   create a Hpa
// @param     hpa               model.Hpa
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateHpa(hpa model.Hpa) (err error) {
	err = global.GVA_DB.Create(&hpa).Error
	return err
}

// @title    DeleteHpa
// @description   delete a Hpa
// @auth                     （2020/04/05  20:22）
// @param     hpa               model.Hpa
// @return                    error

func DeleteHpa(hpa model.Hpa) (err error) {
	err = global.GVA_DB.Delete(hpa).Error
	return err
}

// @title    DeleteHpaByIds
// @description   delete Hpas
// @auth                     （2020/04/05  20:22）
// @param     hpa               model.Hpa
// @return                    error

func DeleteHpaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Hpa{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateHpa
// @description   update a Hpa
// @param     hpa          *model.Hpa
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateHpa(hpa *model.Hpa) (err error) {
	err = global.GVA_DB.Debug().Where("app_id = ? ", hpa.AppId).Updates(hpa).Error
	return err
}

// @title    GetHpa
// @description   get the info of a Hpa
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Hpa        Hpa

func GetHpa(appId uint) (err error, hpa model.Hpa) {
	err = global.GVA_DB.Debug().Where("app_id = ?", appId).First(&hpa).Error
	return
}

// @title    GetHpaInfoList
// @description   get Hpa list by pagination, 分页获取Hpa
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetHpaInfoList(info request.HpaSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Hpa{})
	var hpas []model.Hpa
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&hpas).Error
	return err, hpas, total
}
