package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployRouter(Router *gin.RouterGroup) {
	DeployRouter := Router.Group("deploy").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		DeployRouter.POST("createDeploy", v1.CreateDeploy)             // 新建Deploy
		DeployRouter.DELETE("deleteDeploy", v1.DeleteDeploy)           // 删除Deploy
		DeployRouter.DELETE("deleteDeployByIds", v1.DeleteDeployByIds) // 批量删除Deploy
		DeployRouter.PUT("updateDeploy", v1.UpdateDeploy)              // 更新Deploy
		DeployRouter.GET("findDeploy", v1.FindDeploy)                  // 根据ID获取Deploy
		DeployRouter.GET("findDeployByAppId", v1.FindDeployByAppId)                  // 根据ID获取Deploy
		DeployRouter.GET("getDeployList", v1.GetDeployList)            // 获取Deploy列表
	}
}
