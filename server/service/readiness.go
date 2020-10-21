package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateReadiness
// @description   create a Readiness
// @param     readiness               model.Readiness
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateReadiness(readiness model.Readiness) (err error) {
	err = global.GVA_DB.Create(&readiness).Error
	return err
}

// @title    DeleteReadiness
// @description   delete a Readiness
// @auth                     （2020/04/05  20:22）
// @param     readiness               model.Readiness
// @return                    error

func DeleteReadiness(readiness model.Readiness) (err error) {
	err = global.GVA_DB.Delete(readiness).Error
	return err
}

// @title    DeleteReadinessByIds
// @description   delete Readinesss
// @auth                     （2020/04/05  20:22）
// @param     readiness               model.Readiness
// @return                    error

func DeleteReadinessByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Readiness{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateReadiness
// @description   update a Readiness
// @param     readiness          *model.Readiness
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateReadiness(readiness *model.Readiness) (err error) {
	err = global.GVA_DB.Save(readiness).Error
	return err
}

// @title    GetReadiness
// @description   get the info of a Readiness
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Readiness        Readiness

func GetReadiness(id uint) (err error, readiness model.Readiness) {
	err = global.GVA_DB.Where("id = ?", id).First(&readiness).Error
	return
}

// @title    GetReadinessInfoList
// @description   get Readiness list by pagination, 分页获取Readiness
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetReadinessInfoList(info request.ReadinessSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Readiness{})
	var readinesss []model.Readiness
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&readinesss).Error
	return err, readinesss, total
}
