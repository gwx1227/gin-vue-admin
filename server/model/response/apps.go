package response

import "gin-vue-admin/model"

type AppsSearch struct {
	model.Apps
	Subbusiness
	Business
}


