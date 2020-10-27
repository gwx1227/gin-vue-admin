package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSubBusinessRouter(Router *gin.RouterGroup) {
	SubBusinessRouter := Router.Group("subBusiness").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SubBusinessRouter.POST("createSubBusiness", v1.CreateSubBusiness)             // 新建SubBusiness
		SubBusinessRouter.DELETE("deleteSubBusiness", v1.DeleteSubBusiness)           // 删除SubBusiness
		SubBusinessRouter.DELETE("deleteSubBusinessByIds", v1.DeleteSubBusinessByIds) // 批量删除SubBusiness
		SubBusinessRouter.PUT("updateSubBusiness", v1.UpdateSubBusiness)              // 更新SubBusiness
		SubBusinessRouter.GET("findSubBusiness", v1.FindSubBusiness)                  // 根据ID获取SubBusiness
		SubBusinessRouter.GET("getSubBusinessList", v1.GetSubBusinessList)            // 获取SubBusiness列表
		SubBusinessRouter.GET("getSubBusinessListByBusinessId", v1.GetSubBusinessListByBusinessId)            // 根据一级业务线ID获取SubBusiness列表
	}
}
