// 自动生成模板DeployHistory
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type DeployHistory struct {
	gorm.Model
	AppId      uint    `json:"appId" form:"appId" gorm:"column:app_id;not null;comment:关联应用id"`
	DeployInfo string `json:"deployInfo" form:"deployInfo" gorm:"column:deploy_info;comment:渲染的JOSN数据;type:text;"`
	OpsUser	   string `json:"opsUser" form:"opsUser" gorm:"column:ops_user;comment:'不是操作人';type:varchar(80);not null;size:80;"`
	Result     string `json:"result" form:"result" gorm:"column:result;comment:部署结果"`
	Status     string `json:"status" form:"status" gorm:"column:status;comment:部署状态"`
}

func (DeployHistory) TableName() string {
	return "deploy_history"
}
