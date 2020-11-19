package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateMetrics
// @description   create a Metrics
// @param     metrics               model.Metrics
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateMetrics(metrics model.Metrics) (err error) {
	err = global.GVA_DB.Create(&metrics).Error
	return err
}

// @title    DeleteMetrics
// @description   delete a Metrics
// @auth                     （2020/04/05  20:22）
// @param     metrics               model.Metrics
// @return                    error

func DeleteMetrics(metrics model.Metrics) (err error) {
	err = global.GVA_DB.Delete(metrics).Error
	return err
}

// @title    DeleteMetricsByIds
// @description   delete Metricss
// @auth                     （2020/04/05  20:22）
// @param     metrics               model.Metrics
// @return                    error

func DeleteMetricsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Metrics{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateMetrics
// @description   update a Metrics
// @param     metrics          *model.Metrics
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateMetrics(metrics *model.Metrics) (err error) {
	err = global.GVA_DB.Where("app_id = ?", metrics.AppId).Updates(metrics).Error
	return err
}

// @title    GetMetrics
// @description   get the info of a Metrics
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Metrics        Metrics

func GetMetrics(appId uint) (err error, metrics model.Metrics) {
	err = global.GVA_DB.Where("app_id = ?", appId).First(&metrics).Error
	return
}

// @title    GetMetricsInfoList
// @description   get Metrics list by pagination, 分页获取Metrics
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetMetricsInfoList(info request.MetricsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Metrics{})
	var metricss []model.Metrics
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&metricss).Error
	return err, metricss, total
}
