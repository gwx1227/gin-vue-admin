// 自动生成模板SysConfig
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type SysConfig struct {
	gorm.Model
	Key   string `json:"key" form:"key" gorm:"column:key;comment:配置KEY;not null;default:'CURRENT_CLUSTER';uniqueIndex"`
	Value string `json:"value" form:"value" gorm:"column:value;comment:配置VALUE;not null;default:'gwx_cluster'"`
}

func (SysConfig) TableName() string {
	return "sys_config"
}
