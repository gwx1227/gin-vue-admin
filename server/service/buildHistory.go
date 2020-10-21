package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateBuildHistory
// @description   create a BuildHistory
// @param     buildHistory               model.BuildHistory
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateBuildHistory(buildHistory model.BuildHistory) (err error) {
	err = global.GVA_DB.Create(&buildHistory).Error
	return err
}

// @title    DeleteBuildHistory
// @description   delete a BuildHistory
// @auth                     （2020/04/05  20:22）
// @param     buildHistory               model.BuildHistory
// @return                    error

func DeleteBuildHistory(buildHistory model.BuildHistory) (err error) {
	err = global.GVA_DB.Delete(buildHistory).Error
	return err
}

// @title    DeleteBuildHistoryByIds
// @description   delete BuildHistorys
// @auth                     （2020/04/05  20:22）
// @param     buildHistory               model.BuildHistory
// @return                    error

func DeleteBuildHistoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.BuildHistory{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateBuildHistory
// @description   update a BuildHistory
// @param     buildHistory          *model.BuildHistory
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateBuildHistory(buildHistory *model.BuildHistory) (err error) {
	err = global.GVA_DB.Save(buildHistory).Error
	return err
}

// @title    GetBuildHistory
// @description   get the info of a BuildHistory
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    BuildHistory        BuildHistory

func GetBuildHistory(id uint) (err error, buildHistory model.BuildHistory) {
	err = global.GVA_DB.Where("id = ?", id).First(&buildHistory).Error
	return
}

// @title    GetBuildHistoryInfoList
// @description   get BuildHistory list by pagination, 分页获取BuildHistory
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetBuildHistoryInfoList(info request.BuildHistorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BuildHistory{})
	var buildHistorys []model.BuildHistory
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&buildHistorys).Error
	return err, buildHistorys, total
}
