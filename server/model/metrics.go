// 自动生成模板Metrics
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Metrics struct {
	gorm.Model
	AppId int    `json:"appId" form:"appId" gorm:"column:app_id;comment:关联应用id"`
	Path  string `json:"path" form:"path" gorm:"column:path;comment:监控指标数据获取路径"`
	Port  int    `json:"port" form:"port" gorm:"column:port;comment:监控数据获取端口"`
}

func (Metrics) TableName() string {
	return "metrics"
}
