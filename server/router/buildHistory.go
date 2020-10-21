package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBuildHistoryRouter(Router *gin.RouterGroup) {
	BuildHistoryRouter := Router.Group("buildHistory").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		BuildHistoryRouter.POST("createBuildHistory", v1.CreateBuildHistory)             // 新建BuildHistory
		BuildHistoryRouter.DELETE("deleteBuildHistory", v1.DeleteBuildHistory)           // 删除BuildHistory
		BuildHistoryRouter.DELETE("deleteBuildHistoryByIds", v1.DeleteBuildHistoryByIds) // 批量删除BuildHistory
		BuildHistoryRouter.PUT("updateBuildHistory", v1.UpdateBuildHistory)              // 更新BuildHistory
		BuildHistoryRouter.GET("findBuildHistory", v1.FindBuildHistory)                  // 根据ID获取BuildHistory
		BuildHistoryRouter.GET("getBuildHistoryList", v1.GetBuildHistoryList)            // 获取BuildHistory列表
	}
}
