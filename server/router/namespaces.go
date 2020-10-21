package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitNamespacesRouter(Router *gin.RouterGroup) {
	NamespacesRouter := Router.Group("namespaces").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NamespacesRouter.POST("createNamespaces", v1.CreateNamespaces)             // 新建Namespaces
		NamespacesRouter.DELETE("deleteNamespaces", v1.DeleteNamespaces)           // 删除Namespaces
		NamespacesRouter.DELETE("deleteNamespacesByIds", v1.DeleteNamespacesByIds) // 批量删除Namespaces
		NamespacesRouter.PUT("updateNamespaces", v1.UpdateNamespaces)              // 更新Namespaces
		NamespacesRouter.GET("findNamespaces", v1.FindNamespaces)                  // 根据ID获取Namespaces
		NamespacesRouter.GET("getNamespacesList", v1.GetNamespacesList)            // 获取Namespaces列表
	}
}
