// 自动生成模板BuildHistory
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type BuildHistory struct {
	gorm.Model
	AppId  int    `json:"appId" form:"appId" gorm:"column:app_id;comment:关联应用id"`
	Status string `json:"status" form:"status" gorm:"column:status;comment:构建结果"`
	Tag    string `json:"tag" form:"tag" gorm:"column:tag;comment:构建代码版本"`
}

func (BuildHistory) TableName() string {
	return "build_history"
}
