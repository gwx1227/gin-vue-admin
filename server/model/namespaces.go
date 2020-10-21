// 自动生成模板Namespaces
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Namespaces struct {
	gorm.Model
	Namespace   string `json:"namespace" form:"namespace" gorm:"column:namespace;comment:命名空间;type:varchar(100);size:100;"`
	Create_user string `json:"create_user" form:"create_user" gorm:"column:create_user;comment:创建用户;type:varchar(100);size:100;"`
}

func (Namespaces) TableName() string {
	return "namespaces"
}
