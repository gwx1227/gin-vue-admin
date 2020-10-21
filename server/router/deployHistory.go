package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployHistoryRouter(Router *gin.RouterGroup) {
	DeployHistoryRouter := Router.Group("deployHistory").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DeployHistoryRouter.POST("createDeployHistory", v1.CreateDeployHistory)             // 新建DeployHistory
		DeployHistoryRouter.DELETE("deleteDeployHistory", v1.DeleteDeployHistory)           // 删除DeployHistory
		DeployHistoryRouter.DELETE("deleteDeployHistoryByIds", v1.DeleteDeployHistoryByIds) // 批量删除DeployHistory
		DeployHistoryRouter.PUT("updateDeployHistory", v1.UpdateDeployHistory)              // 更新DeployHistory
		DeployHistoryRouter.GET("findDeployHistory", v1.FindDeployHistory)                  // 根据ID获取DeployHistory
		DeployHistoryRouter.GET("getDeployHistoryList", v1.GetDeployHistoryList)            // 获取DeployHistory列表
	}
}
