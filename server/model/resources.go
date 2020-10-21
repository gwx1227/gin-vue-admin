// 自动生成模板Resources
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Resources struct {
	gorm.Model
	AppId       int    `json:"appId" form:"appId" gorm:"column:app_id;comment:关联应用id"`
	CpuLimit    string `json:"cpuLimit" form:"cpuLimit" gorm:"column:cpu_limit;comment:CPU限额"`
	CpuRequests string `json:"cpuRequests" form:"cpuRequests" gorm:"column:cpu_requests;comment:CPU需求资源"`
	MemLimit    string `json:"memLimit" form:"memLimit" gorm:"column:mem_limit;comment:MEM限额"`
	MemRequests string `json:"memRequests" form:"memRequests" gorm:"column:mem_requests;comment:MEM需求资源"`
}

func (Resources) TableName() string {
	return "resources"
}
