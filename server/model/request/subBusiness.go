package request

import "gin-vue-admin/model"

type SubBusinessSearch struct {
	model.SubBusiness
	PageInfo
}
