package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLivenessRouter(Router *gin.RouterGroup) {
	LivenessRouter := Router.Group("liveness").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		LivenessRouter.POST("createLiveness", v1.CreateLiveness)             // 新建Liveness
		LivenessRouter.DELETE("deleteLiveness", v1.DeleteLiveness)           // 删除Liveness
		LivenessRouter.DELETE("deleteLivenessByIds", v1.DeleteLivenessByIds) // 批量删除Liveness
		LivenessRouter.PUT("updateLiveness", v1.UpdateLiveness)              // 更新Liveness
		LivenessRouter.GET("findLiveness", v1.FindLiveness)                  // 根据ID获取Liveness
		LivenessRouter.GET("getLivenessList", v1.GetLivenessList)            // 获取Liveness列表
	}
}
