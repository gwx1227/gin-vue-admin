// 自动生成模板Hpa
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Hpa struct {
	gorm.Model
	AppId                       int    `json:"appId" form:"appId" gorm:"column:app_id;comment:关联应用id"`
	CpuTargetAverageUtilization int    `json:"cpuTargetAverageUtilization" form:"cpuTargetAverageUtilization" gorm:"column:cpu_target_average_utilization;comment:平均CPU扩展水位线"`
	MaxReplicas                 int    `json:"maxReplicas" form:"maxReplicas" gorm:"column:max_replicas;comment:最大副本数"`
	MemTargetAverageValue       string `json:"memTargetAverageValue" form:"memTargetAverageValue" gorm:"column:mem_target_average_value;comment:平均内存扩展水位线"`
	MinReplicas                 int    `json:"minReplicas" form:"minReplicas" gorm:"column:min_replicas;comment:最小副本数"`
}

func (Hpa) TableName() string {
	return "hpa"
}
