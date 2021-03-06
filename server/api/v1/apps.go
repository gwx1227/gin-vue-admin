package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Apps
// @Summary 创建Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apps true "创建Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apps/createApps [post]
func CreateApps(c *gin.Context) {
	var apps request.AppsSearch
	_ = c.ShouldBindJSON(&apps)
	err,appId := service.CreateApps(model.Apps{AppName: apps.AppName,BusinessId: apps.BusinessId, CreateUser: apps.CreateUser, CurrentEnv: apps.CurrentEnv, GatewaySwitch: apps.GatewaySwitch, GitUrl: apps.GitUrl, HpaSwitch: apps.HpaSwitch,Language: apps.Language, LivenessSwitch: apps.LivenessSwitch, MetricsSwitch: apps.MetricsSwitch, NamespaceId: apps.NamespaceId,ReadinessSwitch: apps.ReadinessSwitch,SubbusinessId: apps.SubbusinessId,UpdateUser: apps.UpdateUser,})
	fmt.Printf("......%v\n\n...%v\n\n", apps.Repository,apps.CommandInfo)
	_ = service.CreateDeploy(model.Deploy{AppId: appId, ArgsSwitch: apps.ArgsSwitch, ArgsInfo: apps.ArgsInfo, CommandInfo: apps.CommandInfo, CommandSwitch: apps.CommandSwitch, ContainerPort: apps.ContainerPort, PullPolicy: apps.PullPolicy, Repository: apps.Repository})
	_ = service.CreateHpa(model.Hpa{AppId: appId})
	_ = service.CreateLiveness(model.Liveness{AppId: appId})
	_ = service.CreateMetrics(model.Metrics{AppId: appId})
	_ = service.CreateReadiness(model.Readiness{AppId: appId})
	_ = service.CreateResources(model.Resources{AppId: appId,CpuLimit: apps.CpuLimit, CpuRequests: apps.CpuRequests, MemLimit: apps.MemLimit, MemRequests: apps.MemRequests})
	if err != nil  {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Apps
// @Summary 删除Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apps true "删除Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apps/deleteApps [delete]
func DeleteApps(c *gin.Context) {
	var apps model.Apps
	_ = c.ShouldBindJSON(&apps)
	err := service.DeleteApps(apps)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Apps
// @Summary 批量删除Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apps/deleteAppsByIds [delete]
func DeleteAppsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteAppsByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Apps
// @Summary 更新Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apps true "更新Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apps/updateApps [put]
func UpdateApps(c *gin.Context) {
	var apps model.Apps
	_ = c.ShouldBindJSON(&apps)
	err := service.UpdateApps(&apps)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func UpdateAppsSwitch(c *gin.Context) {
	var apps model.Apps
	_ = c.ShouldBindJSON(&apps)
	err := service.UpdateAppsSwitch(&apps)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Apps
// @Summary 用id查询Apps
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apps true "用id查询Apps"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apps/findApps [get]
func FindApps(c *gin.Context) {
	var apps model.Apps
	_ = c.ShouldBindQuery(&apps)
	err, reapps := service.GetApps(apps.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reapps": reapps}, c)
	}
}

// @Tags Apps
// @Summary 分页获取Apps列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AppsSearch true "分页获取Apps列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apps/getAppsList [get]
func GetAppsList(c *gin.Context) {
	var pageInfo request.AppsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAppsInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}


func GetAppsListByNamespaceId(c *gin.Context) {
	var pageInfo request.AppsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetAppsInfoListByNamespaceId(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
