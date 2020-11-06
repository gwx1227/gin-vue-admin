package response

import "gin-vue-admin/model"

type AppInfo struct {
	model.Apps
	Subbusiness
	Business
}


