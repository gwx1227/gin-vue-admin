package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateNamespaces
// @description   create a Namespaces
// @param     namespaces               model.Namespaces
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateNamespaces(namespaces model.Namespaces) (err error) {
	err = global.GVA_DB.Create(&namespaces).Error
	return err
}

// @title    DeleteNamespaces
// @description   delete a Namespaces
// @auth                     （2020/04/05  20:22）
// @param     namespaces               model.Namespaces
// @return                    error

func DeleteNamespaces(namespaces model.Namespaces) (err error) {
	err = global.GVA_DB.Delete(namespaces).Error
	return err
}

// @title    DeleteNamespacesByIds
// @description   delete Namespacess
// @auth                     （2020/04/05  20:22）
// @param     namespaces               model.Namespaces
// @return                    error

func DeleteNamespacesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Namespaces{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateNamespaces
// @description   update a Namespaces
// @param     namespaces          *model.Namespaces
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateNamespaces(namespaces *model.Namespaces) (err error) {
	err = global.GVA_DB.Save(namespaces).Error
	return err
}

// @title    GetNamespaces
// @description   get the info of a Namespaces
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Namespaces        Namespaces

func GetNamespaces(id uint) (err error, namespaces model.Namespaces) {
	fmt.Printf("查询ID: %d", id)
	err = global.GVA_DB.Where("id = ?", id).First(&namespaces).Error
	return
}

// @title    GetNamespacesInfoList
// @description   get Namespaces list by pagination, 分页获取Namespaces
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetNamespacesInfoList(info request.NamespacesSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Namespaces{})
	var namespacess []model.Namespaces
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Namespace != "" {
		db = db.Where("`namespace` = ?", info.Namespace)
	}
	if info.Create_user != "" {
		db = db.Where("`create_user` = ?", info.Create_user)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&namespacess).Error
	return err, namespacess, total
}
