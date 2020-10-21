// 自动生成模板Gateway
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Gateway struct {
	gorm.Model
	AppId        int    `json:"appId" form:"appId" gorm:"column:app_id;comment:关联应用id"`
	Hosts        string `json:"hosts" form:"hosts" gorm:"column:hosts;comment:域名信息"`
	Paths        string `json:"paths" form:"paths" gorm:"column:paths;comment:域名对应启用路径"`
	WeightCanary int    `json:"weightCanary" form:"weightCanary" gorm:"column:weight_canary;comment:canary流量权重"`
	WeightOnline int    `json:"weightOnline" form:"weightOnline" gorm:"column:weight_online;comment:online流量权重"`
}

func (Gateway) TableName() string {
	return "gateway"
}
