package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAppsRouter(Router *gin.RouterGroup) {
	AppsRouter := Router.Group("apps").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AppsRouter.POST("createApps", v1.CreateApps)             // 新建Apps
		AppsRouter.DELETE("deleteApps", v1.DeleteApps)           // 删除Apps
		AppsRouter.DELETE("deleteAppsByIds", v1.DeleteAppsByIds) // 批量删除Apps
		AppsRouter.PUT("updateApps", v1.UpdateApps)              // 更新Apps
		AppsRouter.GET("findApps", v1.FindApps)                  // 根据ID获取Apps
		AppsRouter.GET("getAppsList", v1.GetAppsList)            // 获取Apps列表
		AppsRouter.GET("getAppsListByNamespaceId", v1.GetAppsListByNamespaceId)            // 根据命令空间ID获取Apps列表
	}
}
