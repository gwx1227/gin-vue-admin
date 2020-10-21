package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitResourcesRouter(Router *gin.RouterGroup) {
	ResourcesRouter := Router.Group("resources").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ResourcesRouter.POST("createResources", v1.CreateResources)             // 新建Resources
		ResourcesRouter.DELETE("deleteResources", v1.DeleteResources)           // 删除Resources
		ResourcesRouter.DELETE("deleteResourcesByIds", v1.DeleteResourcesByIds) // 批量删除Resources
		ResourcesRouter.PUT("updateResources", v1.UpdateResources)              // 更新Resources
		ResourcesRouter.GET("findResources", v1.FindResources)                  // 根据ID获取Resources
		ResourcesRouter.GET("getResourcesList", v1.GetResourcesList)            // 获取Resources列表
	}
}
