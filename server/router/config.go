package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigRouter := Router.Group("config").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ConfigRouter.POST("createConfig", v1.CreateConfig)             // 新建Config
		ConfigRouter.DELETE("deleteConfig", v1.DeleteConfig)           // 删除Config
		ConfigRouter.DELETE("deleteConfigByIds", v1.DeleteConfigByIds) // 批量删除Config
		ConfigRouter.PUT("updateConfig", v1.UpdateConfig)              // 更新Config
		ConfigRouter.GET("findConfig", v1.FindConfig)                  // 根据ID获取Config
		ConfigRouter.GET("getConfigList", v1.GetConfigList)            // 获取Config列表
		ConfigRouter.GET("getConfigListByAppId", v1.GetConfigListByAppId)            // 根据应用ID获取Config列表
	}
}
