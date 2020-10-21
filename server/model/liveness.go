// 自动生成模板Liveness
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Liveness struct {
	gorm.Model
	AppId               int    `json:"appId" form:"appId" gorm:"column:app_id;comment:关联应用id"`
	FailureThreshold    int    `json:"failureThreshold" form:"failureThreshold" gorm:"column:failure_threshold;comment:失败次数"`
	InitialDelaySeconds int    `json:"initialDelaySeconds" form:"initialDelaySeconds" gorm:"column:initial_delay_seconds;comment:初始化延迟探测时间"`
	Path                string `json:"path" form:"path" gorm:"column:path;comment:健康检查路径"`
	PeriodSeconds       int    `json:"periodSeconds" form:"periodSeconds" gorm:"column:period_seconds;comment:探测间隔时间"`
	SuccessThreshold    int    `json:"successThreshold" form:"successThreshold" gorm:"column:success_threshold;comment:探测成功次数"`
	TimeoutSeconds      int    `json:"timeoutSeconds" form:"timeoutSeconds" gorm:"column:timeout_seconds;comment:探测超时时间"`
}

func (Liveness) TableName() string {
	return "liveness"
}
