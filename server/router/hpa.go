package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitHpaRouter(Router *gin.RouterGroup) {
	HpaRouter := Router.Group("hpa").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		HpaRouter.POST("createHpa", v1.CreateHpa)             // 新建Hpa
		HpaRouter.DELETE("deleteHpa", v1.DeleteHpa)           // 删除Hpa
		HpaRouter.DELETE("deleteHpaByIds", v1.DeleteHpaByIds) // 批量删除Hpa
		HpaRouter.PUT("updateHpa", v1.UpdateHpa)              // 更新Hpa
		HpaRouter.GET("findHpa", v1.FindHpa)                  // 根据ID获取Hpa
		HpaRouter.GET("getHpaList", v1.GetHpaList)            // 获取Hpa列表
	}
}
