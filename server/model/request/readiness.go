package request

import "gin-vue-admin/model"

type ReadinessSearch struct {
	model.Readiness
	PageInfo
}
