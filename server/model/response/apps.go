package response

import "gin-vue-admin/model"

type AppInfo struct {
	model.Apps
	Subbusiness
	Business
}

type AppInfoNamespacs struct {
	model.Apps
	Namespace   string `json:"namespace"`
}


