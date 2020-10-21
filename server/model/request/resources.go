package request

import "gin-vue-admin/model"

type ResourcesSearch struct {
	model.Resources
	PageInfo
}
