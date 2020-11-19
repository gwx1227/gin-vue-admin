package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateResources
// @description   create a Resources
// @param     resources               model.Resources
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateResources(resources model.Resources) (err error) {
	err = global.GVA_DB.Create(&resources).Error
	return err
}

// @title    DeleteResources
// @description   delete a Resources
// @auth                     （2020/04/05  20:22）
// @param     resources               model.Resources
// @return                    error

func DeleteResources(resources model.Resources) (err error) {
	err = global.GVA_DB.Delete(resources).Error
	return err
}

// @title    DeleteResourcesByIds
// @description   delete Resourcess
// @auth                     （2020/04/05  20:22）
// @param     resources               model.Resources
// @return                    error

func DeleteResourcesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Resources{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateResources
// @description   update a Resources
// @param     resources          *model.Resources
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateResources(resources *model.Resources) (err error) {
	err = global.GVA_DB.Where("app_id = ?", resources.AppId).Updates(resources).Error
	return err
}

// @title    GetResources
// @description   get the info of a Resources
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Resources        Resources

func GetResources(appId uint) (err error, resources model.Resources) {
	err = global.GVA_DB.Where("app_id = ?", appId).First(&resources).Error
	return
}
func GetResourcesByAppId(appId uint) (err error, resources model.Resources) {
	err = global.GVA_DB.Where("app_id = ?", appId).First(&resources).Error
	return
}
// @title    GetResourcesInfoList
// @description   get Resources list by pagination, 分页获取Resources
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetResourcesInfoList(info request.ResourcesSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Resources{})
	var resourcess []model.Resources
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&resourcess).Error
	return err, resourcess, total
}
