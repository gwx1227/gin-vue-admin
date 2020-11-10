package request

import "gin-vue-admin/model"

type AppsSearch struct {
	model.Apps
	CreateApp
	PageInfo
}
