package request

import "gin-vue-admin/model"

type HpaSearch struct {
	model.Hpa
	PageInfo
}
