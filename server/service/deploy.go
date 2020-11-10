package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateDeploy
// @description   create a Deploy
// @param     deploy               model.Deploy
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateDeploy(deploy model.Deploy) (err error) {
	err = global.GVA_DB.Debug().Create(&deploy).Error
	return err
}

// @title    DeleteDeploy
// @description   delete a Deploy
// @auth                     （2020/04/05  20:22）
// @param     deploy               model.Deploy
// @return                    error

func DeleteDeploy(deploy model.Deploy) (err error) {
	err = global.GVA_DB.Delete(deploy).Error
	return err
}

// @title    DeleteDeployByIds
// @description   delete Deploys
// @auth                     （2020/04/05  20:22）
// @param     deploy               model.Deploy
// @return                    error

func DeleteDeployByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Deploy{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateDeploy
// @description   update a Deploy
// @param     deploy          *model.Deploy
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateDeploy(deploy *model.Deploy) (err error) {
	err = global.GVA_DB.Debug().Where("app_id = ?", deploy.AppId).Updates(deploy).Error
	return err
}

// @title    GetDeploy
// @description   get the info of a Deploy
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Deploy        Deploy

func GetDeploy(id uint) (err error, deploy model.Deploy) {
	err = global.GVA_DB.Where("id = ?", id).First(&deploy).Error
	return
}

func GetDeployByAppId(app_id uint) (err error, deploy model.Deploy) {
	err = global.GVA_DB.Debug().Where("app_id = ?", app_id).First(&deploy).Error
	return
}
// @title    GetDeployInfoList
// @description   get Deploy list by pagination, 分页获取Deploy
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetDeployInfoList(info request.DeploySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Deploy{})
	var deploys []model.Deploy
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&deploys).Error
	return err, deploys, total
}
