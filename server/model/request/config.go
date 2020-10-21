package request

import "gin-vue-admin/model"

type ConfigSearch struct {
	model.Config
	PageInfo
}
