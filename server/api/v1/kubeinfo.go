package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

func GetPodList(c *gin.Context) {
	var pageInfo request.AppsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	_,appInfo := service.GetAppsByAppId(pageInfo.ID)
	podList,err := service.GetPodList(appInfo.Namespace, appInfo.AppName,c.Query("pod_env"))
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     podList,
		}, c)
	}
}

func RestartPod(c *gin.Context) {
	var pageInfo request.AppsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	_,appInfo := service.GetAppsByAppId(pageInfo.ID)
	res,err := service.DeletePod(appInfo.Namespace,c.PostForm("pod_name"))
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(res,c)
	}
}

func GetDeployment(c *gin.Context)  {
	var pageInfo request.AppsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	_,appInfo := service.GetAppsByAppId(pageInfo.ID)
	deployment,err := service.Getdeployment(appInfo.Namespace, appInfo.AppName, c.Query("deployment_name"))
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
				List:     deployment,
		}, c)
	}
}

func GetDeploymentList(c *gin.Context)  {
	var pageInfo request.AppsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	_,appInfo := service.GetAppsByAppId(pageInfo.ID)
	deploymentList,err := service.GetdeploymentList(appInfo.Namespace,appInfo.AppName)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     deploymentList,
		}, c)
	}
}

