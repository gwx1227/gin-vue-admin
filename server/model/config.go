// 自动生成模板Config
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Config struct {
	gorm.Model
	AppId int    `json:"appId" form:"appId" gorm:"column:app_id;comment:关联app_id"`
	Key   string `json:"key" form:"key" gorm:"column:key;comment:配置KEY"`
	Value string `json:"value" form:"value" gorm:"column:value;comment:配置VALUE"`
}

func (Config) TableName() string {
	return "config"
}
