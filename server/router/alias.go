package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAliasRouter(Router *gin.RouterGroup) {
	AliasRouter := Router.Group("alias").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AliasRouter.POST("createAlias", v1.CreateAlias)             // 新建Alias
		AliasRouter.DELETE("deleteAlias", v1.DeleteAlias)           // 删除Alias
		AliasRouter.DELETE("deleteAliasByIds", v1.DeleteAliasByIds) // 批量删除Alias
		AliasRouter.PUT("updateAlias", v1.UpdateAlias)              // 更新Alias
		AliasRouter.GET("findAlias", v1.FindAlias)                  // 根据ID获取Alias
		AliasRouter.GET("getAliasList", v1.GetAliasList)            // 获取Alias列表
		AliasRouter.GET("getAliasListByAppId", v1.GetAliasListByAppId)            // 根据应用ID获取Alias列表
	}
}
