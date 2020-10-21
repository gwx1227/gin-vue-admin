package request

import "gin-vue-admin/model"

type AliasSearch struct {
	model.Alias
	PageInfo
}
