// 自动生成模板Alias
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Alias struct {
	gorm.Model
	AppId    int    `json:"appId" form:"appId" gorm:"column:app_id;index;not null;uniqueIndex:uix_app_hostname_app_id_hostname;comment:'关联应用id';type:int(10);size:10;"`
	Hostname string `json:"hostname" form:"hostname" gorm:"column:hostname;uniqueIndex:uix_app_hostname_app_id_hostname;comment:'劫持域名';type:varchar(80);size:80;"`
	Ip       string `json:"ip" form:"ip" gorm:"column:ip;comment:'劫持IP';type:varchar(30);size:30;"`
}

func (Alias) TableName() string {
	return "alias"
}
