package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitGatewayRouter(Router *gin.RouterGroup) {
	GatewayRouter := Router.Group("gateway").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		GatewayRouter.POST("createGateway", v1.CreateGateway)             // 新建Gateway
		GatewayRouter.DELETE("deleteGateway", v1.DeleteGateway)           // 删除Gateway
		GatewayRouter.DELETE("deleteGatewayByIds", v1.DeleteGatewayByIds) // 批量删除Gateway
		GatewayRouter.PUT("updateGateway", v1.UpdateGateway)              // 更新Gateway
		GatewayRouter.GET("findGateway", v1.FindGateway)                  // 根据ID获取Gateway
		GatewayRouter.GET("getGatewayList", v1.GetGatewayList)            // 获取Gateway列表
	}
}
