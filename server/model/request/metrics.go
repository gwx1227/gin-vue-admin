package request

import "gin-vue-admin/model"

type MetricsSearch struct {
	model.Metrics
	PageInfo
}
