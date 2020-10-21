package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateDeployHistory
// @description   create a DeployHistory
// @param     deployHistory               model.DeployHistory
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateDeployHistory(deployHistory model.DeployHistory) (err error) {
	err = global.GVA_DB.Create(&deployHistory).Error
	return err
}

// @title    DeleteDeployHistory
// @description   delete a DeployHistory
// @auth                     （2020/04/05  20:22）
// @param     deployHistory               model.DeployHistory
// @return                    error

func DeleteDeployHistory(deployHistory model.DeployHistory) (err error) {
	err = global.GVA_DB.Delete(deployHistory).Error
	return err
}

// @title    DeleteDeployHistoryByIds
// @description   delete DeployHistorys
// @auth                     （2020/04/05  20:22）
// @param     deployHistory               model.DeployHistory
// @return                    error

func DeleteDeployHistoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DeployHistory{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateDeployHistory
// @description   update a DeployHistory
// @param     deployHistory          *model.DeployHistory
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateDeployHistory(deployHistory *model.DeployHistory) (err error) {
	err = global.GVA_DB.Save(deployHistory).Error
	return err
}

// @title    GetDeployHistory
// @description   get the info of a DeployHistory
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    DeployHistory        DeployHistory

func GetDeployHistory(id uint) (err error, deployHistory model.DeployHistory) {
	err = global.GVA_DB.Where("id = ?", id).First(&deployHistory).Error
	return
}

// @title    GetDeployHistoryInfoList
// @description   get DeployHistory list by pagination, 分页获取DeployHistory
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetDeployHistoryInfoList(info request.DeployHistorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.DeployHistory{})
	var deployHistorys []model.DeployHistory
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&deployHistorys).Error
	return err, deployHistorys, total
}
