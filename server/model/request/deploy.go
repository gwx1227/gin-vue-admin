package request

import "gin-vue-admin/model"

type DeploySearch struct {
	model.Deploy
	PageInfo
}
