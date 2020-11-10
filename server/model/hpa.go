// 自动生成模板Hpa
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Hpa struct {
	gorm.Model
	AppId                       uint    `json:"appId" form:"appId" gorm:"column:app_id;uniqueIndex;not null;comment:关联应用id"`
	CpuTargetAverageUtilization int    `json:"cpuTargetAverageUtilization" form:"cpuTargetAverageUtilization" gorm:"column:cpu_target_average_utilization;not null;default:60;comment:平均CPU扩展水位线"`
	MaxReplicas                 int    `json:"maxReplicas" form:"maxReplicas" gorm:"column:max_replicas;not null;default:20;comment:最大副本数"`
	MemTargetAverageValue       string `json:"memTargetAverageValue" form:"memTargetAverageValue" gorm:"column:mem_target_average_value;not null;default:'3Gi';comment:平均内存扩展水位线"`
	MinReplicas                 int    `json:"minReplicas" form:"minReplicas" gorm:"column:min_replicas;not null;default:2;comment:最小副本数"`
}

func (Hpa) TableName() string {
	return "hpa"
}
