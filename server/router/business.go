package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBusinessRouter(Router *gin.RouterGroup) {
	BusinessRouter := Router.Group("business").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		BusinessRouter.POST("createBusiness", v1.CreateBusiness)             // 新建Business
		BusinessRouter.DELETE("deleteBusiness", v1.DeleteBusiness)           // 删除Business
		BusinessRouter.DELETE("deleteBusinessByIds", v1.DeleteBusinessByIds) // 批量删除Business
		BusinessRouter.PUT("updateBusiness", v1.UpdateBusiness)              // 更新Business
		BusinessRouter.GET("findBusiness", v1.FindBusiness)                  // 根据ID获取Business
		BusinessRouter.GET("getBusinessList", v1.GetBusinessList)            // 获取Business列表
	}
}
