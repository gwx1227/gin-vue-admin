package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysConfigRouter(Router *gin.RouterGroup) {
	SysConfigRouter := Router.Group("sysConfig").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SysConfigRouter.POST("createSysConfig", v1.CreateSysConfig)             // 新建SysConfig
		SysConfigRouter.DELETE("deleteSysConfig", v1.DeleteSysConfig)           // 删除SysConfig
		SysConfigRouter.DELETE("deleteSysConfigByIds", v1.DeleteSysConfigByIds) // 批量删除SysConfig
		SysConfigRouter.PUT("updateSysConfig", v1.UpdateSysConfig)              // 更新SysConfig
		SysConfigRouter.GET("findSysConfig", v1.FindSysConfig)                  // 根据ID获取SysConfig
		SysConfigRouter.GET("getSysConfigList", v1.GetSysConfigList)            // 获取SysConfig列表
	}
}
