package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateLiveness
// @description   create a Liveness
// @param     liveness               model.Liveness
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLiveness(liveness model.Liveness) (err error) {
	err = global.GVA_DB.Create(&liveness).Error
	return err
}

// @title    DeleteLiveness
// @description   delete a Liveness
// @auth                     （2020/04/05  20:22）
// @param     liveness               model.Liveness
// @return                    error

func DeleteLiveness(liveness model.Liveness) (err error) {
	err = global.GVA_DB.Delete(liveness).Error
	return err
}

// @title    DeleteLivenessByIds
// @description   delete Livenesss
// @auth                     （2020/04/05  20:22）
// @param     liveness               model.Liveness
// @return                    error

func DeleteLivenessByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Liveness{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateLiveness
// @description   update a Liveness
// @param     liveness          *model.Liveness
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLiveness(liveness *model.Liveness) (err error) {
	err = global.GVA_DB.Debug().Where("app_id = ? ", liveness.AppId).Updates(liveness).Error
	return err
}

// @title    GetLiveness
// @description   get the info of a Liveness
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Liveness        Liveness

func GetLiveness(appId uint) (err error, liveness model.Liveness) {
	err = global.GVA_DB.Debug().Where("app_id = ?", appId).First(&liveness).Error
	return
}

// @title    GetLivenessInfoList
// @description   get Liveness list by pagination, 分页获取Liveness
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLivenessInfoList(info request.LivenessSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Liveness{})
	var livenesss []model.Liveness
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&livenesss).Error
	return err, livenesss, total
}
