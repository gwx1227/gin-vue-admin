package request

import "gin-vue-admin/model"

type LivenessSearch struct {
	model.Liveness
	PageInfo
}
