package request

import "gin-vue-admin/model"

type BuildHistorySearch struct {
	model.BuildHistory
	PageInfo
}
