package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateBusiness
// @description   create a Business
// @param     business               model.Businesss
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateBusiness(business model.Businesss) (err error) {
	err = global.GVA_DB.Create(&business).Error
	return err
}

// @title    DeleteBusiness
// @description   delete a Business
// @auth                     （2020/04/05  20:22）
// @param     business               model.Businesss
// @return                    error

func DeleteBusiness(business model.Businesss) (err error) {
	err = global.GVA_DB.Delete(business).Error
	return err
}

// @title    DeleteBusinessByIds
// @description   delete Businesss
// @auth                     （2020/04/05  20:22）
// @param     business               model.Businesss
// @return                    error

func DeleteBusinessByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Businesss{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateBusiness
// @description   update a Business
// @param     business          *model.Businesss
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateBusiness(business *model.Businesss) (err error) {
	err = global.GVA_DB.Save(business).Error
	return err
}

// @title    GetBusiness
// @description   get the info of a Business
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Business        Business

func GetBusiness(id uint) (err error, business model.Businesss) {
	err = global.GVA_DB.Where("id = ?", id).First(&business).Error
	return
}

// @title    GetBusinessInfoList
// @description   get Business list by pagination, 分页获取Business
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetBusinessInfoList(info request.BusinessSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Businesss{})
	var businesss []model.Businesss
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Business != "" {
		db = db.Where("`business` = ?", info.Business)
	}
	if info.Business_owner != "" {
		db = db.Where("`business_owner` = ?", info.Business_owner)
	}
	if info.Business_sre != "" {
		db = db.Where("`business_sre` = ?", info.Business_sre)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&businesss).Error
	return err, businesss, total
}
