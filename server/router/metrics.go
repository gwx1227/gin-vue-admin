package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMetricsRouter(Router *gin.RouterGroup) {
	MetricsRouter := Router.Group("metrics").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		MetricsRouter.POST("createMetrics", v1.CreateMetrics)             // 新建Metrics
		MetricsRouter.DELETE("deleteMetrics", v1.DeleteMetrics)           // 删除Metrics
		MetricsRouter.DELETE("deleteMetricsByIds", v1.DeleteMetricsByIds) // 批量删除Metrics
		MetricsRouter.PUT("updateMetrics", v1.UpdateMetrics)              // 更新Metrics
		MetricsRouter.GET("findMetrics", v1.FindMetrics)                  // 根据ID获取Metrics
		MetricsRouter.GET("getMetricsList", v1.GetMetricsList)            // 获取Metrics列表
	}
}
