package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitReadinessRouter(Router *gin.RouterGroup) {
	ReadinessRouter := Router.Group("readiness").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ReadinessRouter.POST("createReadiness", v1.CreateReadiness)             // 新建Readiness
		ReadinessRouter.DELETE("deleteReadiness", v1.DeleteReadiness)           // 删除Readiness
		ReadinessRouter.DELETE("deleteReadinessByIds", v1.DeleteReadinessByIds) // 批量删除Readiness
		ReadinessRouter.PUT("updateReadiness", v1.UpdateReadiness)              // 更新Readiness
		ReadinessRouter.GET("findReadiness", v1.FindReadiness)                  // 根据ID获取Readiness
		ReadinessRouter.GET("getReadinessList", v1.GetReadinessList)            // 获取Readiness列表
	}
}
