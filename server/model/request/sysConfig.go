package request

import "gin-vue-admin/model"

type SysConfigSearch struct {
	model.SysConfig
	PageInfo
}
