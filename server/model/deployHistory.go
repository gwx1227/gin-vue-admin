// 自动生成模板DeployHistory
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type DeployHistory struct {
	gorm.Model
	AppId      int    `json:"appId" form:"appId" gorm:"column:app_id;comment:关联应用id"`
	DeployInfo string `json:"deployInfo" form:"deployInfo" gorm:"column:deploy_info;comment:渲染的JOSN数据;type:text;"`
	Env        string `json:"env" form:"env" gorm:"column:env;comment:环境"`
	Result     string `json:"result" form:"result" gorm:"column:result;comment:部署结果"`
	Status     string `json:"status" form:"status" gorm:"column:status;comment:部署状态"`
}

func (DeployHistory) TableName() string {
	return "deploy_history"
}
