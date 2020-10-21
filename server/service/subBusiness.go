package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSubBusiness
// @description   create a SubBusiness
// @param     subBusiness               model.SubBusiness
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSubBusiness(subBusiness model.SubBusiness) (err error) {
	err = global.GVA_DB.Create(&subBusiness).Error
	return err
}

// @title    DeleteSubBusiness
// @description   delete a SubBusiness
// @auth                     （2020/04/05  20:22）
// @param     subBusiness               model.SubBusiness
// @return                    error

func DeleteSubBusiness(subBusiness model.SubBusiness) (err error) {
	err = global.GVA_DB.Delete(subBusiness).Error
	return err
}

// @title    DeleteSubBusinessByIds
// @description   delete SubBusinesss
// @auth                     （2020/04/05  20:22）
// @param     subBusiness               model.SubBusiness
// @return                    error

func DeleteSubBusinessByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SubBusiness{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateSubBusiness
// @description   update a SubBusiness
// @param     subBusiness          *model.SubBusiness
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSubBusiness(subBusiness *model.SubBusiness) (err error) {
	err = global.GVA_DB.Save(subBusiness).Error
	return err
}

// @title    GetSubBusiness
// @description   get the info of a SubBusiness
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    SubBusiness        SubBusiness

func GetSubBusiness(id uint) (err error, subBusiness model.SubBusiness) {
	err = global.GVA_DB.Where("id = ?", id).First(&subBusiness).Error
	return
}

// @title    GetSubBusinessInfoList
// @description   get SubBusiness list by pagination, 分页获取SubBusiness
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSubBusinessInfoList(info request.SubBusinessSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SubBusiness{})
	var subBusinesss []model.SubBusiness
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&subBusinesss).Error
	return err, subBusinesss, total
}
