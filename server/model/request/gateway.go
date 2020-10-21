package request

import "gin-vue-admin/model"

type GatewaySearch struct {
	model.Gateway
	PageInfo
}
