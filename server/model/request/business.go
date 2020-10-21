package request

import "gin-vue-admin/model"

type BusinessSearch struct {
	model.Businesss
	PageInfo
}
