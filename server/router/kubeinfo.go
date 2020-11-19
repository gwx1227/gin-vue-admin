package router

import (
"gin-vue-admin/api/v1"
"gin-vue-admin/middleware"
"github.com/gin-gonic/gin"
)

func InitKubeRouter(Router *gin.RouterGroup) {
	AppsRouter := Router.Group("kubeinfo").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		AppsRouter.POST("restartPod", v1.RestartPod)             // 获取应用pod列表
		AppsRouter.GET("getPodList", v1.GetPodList)             // 获取应用pod列表
		AppsRouter.GET("getDeployment", v1.GetDeployment)             // 获取应用Deployment信息
		AppsRouter.GET("getDeploymentList", v1.GetDeploymentList)             // 获取应用Deployment列表

	}
}



