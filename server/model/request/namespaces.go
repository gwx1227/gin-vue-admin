package request

import "gin-vue-admin/model"

type NamespacesSearch struct {
	model.Namespaces
	PageInfo
}
