// 自动生成模板SubBusiness
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type SubBusiness struct {
	gorm.Model
	BusinessId       int    `json:"businessId" form:"businessId" gorm:"column:business_id;comment:父业务线ID"`
	Subbusiness      string `json:"subbusiness" form:"subbusiness" gorm:"column:subbusiness;comment:二级业务线"`
	SubbusinessOwner string `json:"subbusinessOwner" form:"subbusinessOwner" gorm:"column:subbusiness_owner;comment:二级业务线负责人"`
}

func (SubBusiness) TableName() string {
	return "sub_business"
}
