// 自动生成模板Business
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Businesss struct {
	gorm.Model
	Business       string `json:"business" form:"business" gorm:"column:business;comment:一级业务线名称;type:varchar(100);size:100;index;not null;unique"`
	Business_owner string `json:"business_owner" form:"business_owner" gorm:"column:business_owner;comment:一级业务线负责人;type:varchar(100);size:100;not null"`
	Business_sre   string `json:"business_sre" form:"business_sre" gorm:"column:business_sre;comment:一级业务线SRE;type:varchar(100);size:100;not null"`
}

func (Businesss) TableName() string {
	return "business"
}
