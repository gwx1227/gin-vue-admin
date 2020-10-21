package request

import "gin-vue-admin/model"

type DeployHistorySearch struct {
	model.DeployHistory
	PageInfo
}
