// 自动生成模板Apps
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Apps struct {
	gorm.Model
	AppName         string `json:"appName" form:"appName" gorm:"column:app_name;comment:'应用名称';type:varchar(100);size:100;"`
	BusinessId      int    `json:"businessId" form:"businessId" gorm:"column:business_id;comment:;type:int(10);size:10;"`
	CreateUser      string `json:"createUser" form:"createUser" gorm:"column:create_user;comment:'创建人';type:varchar(80);size:80;"`
	CurrentEnv      string `json:"currentEnv" form:"currentEnv" gorm:"column:current_env;comment:'当前应用环境';type:varchar(30);size:30;"`
	GatewaySwitch   *bool  `json:"gatewaySwitch" form:"gatewaySwitch" gorm:"column:gateway_switch;comment:'域名开关';type:tinyint;"`
	GitUrl          string `json:"gitUrl" form:"gitUrl" gorm:"column:git_url;comment:'代码地址';type:varchar(200);size:200;"`
	HpaSwitch       *bool  `json:"hpaSwitch" form:"hpaSwitch" gorm:"column:hpa_switch;comment:'自动扩展开关';type:tinyint;"`
	Language        string `json:"language" form:"language" gorm:"column:language;comment:'应用语言';type:varchar(80);size:80;"`
	LivenessSwitch  *bool  `json:"livenessSwitch" form:"livenessSwitch" gorm:"column:liveness_switch;comment:'健康检查开关';type:tinyint;"`
	MetricsSwitch   *bool  `json:"metricsSwitch" form:"metricsSwitch" gorm:"column:metrics_switch;comment:'监控接入开关';type:tinyint;"`
	NamespaceId     int    `json:"namespaceId" form:"namespaceId" gorm:"column:namespace_id;comment:;type:int(10);size:10;"`
	ReadinessSwitch *bool  `json:"readinessSwitch" form:"readinessSwitch" gorm:"column:readiness_switch;comment:'就绪检查开关';type:tinyint;"`
	SubbusinessId   int    `json:"subbusinessId" form:"subbusinessId" gorm:"column:subbusiness_id;comment:;type:int(10);size:10;"`
	UpdateUser      string `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:'更新人';type:varchar(40);size:40;"`
	Usable          *bool  `json:"usable" form:"usable" gorm:"column:usable;comment:'服役状态标志位,代替删除';type:tinyint;"`
}

func (Apps) TableName() string {
	return "apps"
}
