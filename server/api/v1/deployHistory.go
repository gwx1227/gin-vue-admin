package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"github.com/lizongshen/gocommand"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)







//
func GetDeployInfo (appId uint)	(deployInfo string)   {
	_,aliasData := service.GetAliasByAppId(appId)
	_,gatewayData := service.GetGatewayByAppId(appId)
	_,gatewayHosts := service.GetGatewayHostByAppId(appId)
	_,appDeployInfo := service.GetDeployInfoByAppId(appId)
	_,switchInfo  := service.GetDeploySwitchByAppId(appId)
	_,livenessInfo := service.GetLivenessByAppId(appId)
	_,readinessInfo := service.GetReadinessByAppId(appId)
	_,resourcesInfo := service.GetResourcesByAppId(appId)

	var data = ``

	t := resp.DeployInfo{}
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	t.Metrics.Enabled = switchInfo.MetricsSwitch
	t.Gateway.Enabled = switchInfo.GatewaySwitch
	t.Liveness.Enabled = switchInfo.LivenessSwitch
	t.Readiness.Enabled = switchInfo.ReadinessSwitch
	t.HpaMode.Enabled = switchInfo.HpaSwitch
	t.Image.ArgsEnable = switchInfo.ArgsSwitch
	t.Image.CommandEnable = switchInfo.CommandSwitch
	t.AppName = appDeployInfo.AppName
	t.Namespace = appDeployInfo.Namespace
	t.CurrentEnv = appDeployInfo.CurrentEnv
	t.Resources.Limits.Cpu = resourcesInfo.CpuLimit
	t.Resources.Limits.Memory = resourcesInfo.MemLimit
	t.Resources.Requests.Cpu = resourcesInfo.CpuRequests
	t.Resources.Requests.Memory = resourcesInfo.MemRequests
	t.Image.WeightCanary = appDeployInfo.Image.WeightCanary
	t.Image.WeightOnline = appDeployInfo.Image.WeightOnline
	t.Image.ArgsInfo = appDeployInfo.Image.ArgsInfo
	t.Image.CommandInfo = appDeployInfo.Image.CommandInfo
	t.Image.Repository = appDeployInfo.Image.Repository
	t.Image.ContainerPort = appDeployInfo.Image.ContainerPort
	t.Image.PullPolicy = appDeployInfo.Image.PullPolicy
	t.Image.ReplicaCountCanary = appDeployInfo.Image.ReplicaCountCanary
	t.Image.ReplicaCountOnline = appDeployInfo.Image.ReplicaCountOnline
	t.Image.TagCanary = appDeployInfo.Image.TagCanary
	t.Image.TagOnline = appDeployInfo.Image.TagOnline
	t.Metrics.Port = appDeployInfo.Metrics.Port
	t.Metrics.Path = appDeployInfo.Metrics.Path
	t.HpaMode.CpuTargetAverageUtilization = appDeployInfo.HpaMode.CpuTargetAverageUtilization
	t.HpaMode.MemTargetAverageValue = appDeployInfo.HpaMode.MemTargetAverageValue
	t.HpaMode.MaxReplicas = appDeployInfo.HpaMode.MaxReplicas
	t.HpaMode.MinReplicas = appDeployInfo.HpaMode.MinReplicas
	t.Liveness.Path = livenessInfo.Path
	t.Liveness.FailureThreshold = livenessInfo.FailureThreshold
	t.Liveness.InitialDelaySeconds = livenessInfo.InitialDelaySeconds
	t.Liveness.PeriodSeconds = livenessInfo.PeriodSeconds
	t.Liveness.SuccessThreshold = livenessInfo.SuccessThreshold
	t.Liveness.TimeoutSeconds = livenessInfo.TimeoutSeconds
	t.Readiness.Path = readinessInfo.Path
	t.Readiness.FailureThreshold = readinessInfo.FailureThreshold
	t.Readiness.InitialDelaySeconds = readinessInfo.InitialDelaySeconds
	t.Readiness.PeriodSeconds = readinessInfo.PeriodSeconds
	t.Readiness.SuccessThreshold = readinessInfo.SuccessThreshold
	t.Readiness.TimeoutSeconds = readinessInfo.TimeoutSeconds
	t.HostAlias.Enabled = true
	t.Service.ServiceType = "ClusterIP"
	t.Service.Port = 80
	for g_index := range gatewayData {
		t.GatewayData = append(t.GatewayData, gatewayData[g_index])
	}
	for h_index := range gatewayHosts {
		t.GatewayHosts = append(t.GatewayHosts, gatewayHosts[h_index])
	}
	for a_index := range aliasData {
		t.PodHostAlias = append(t.PodHostAlias, aliasData[a_index])
	}
	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- t dump:\n%s\n\n", string(d))
	return string(d)
}


// @Tags DeployHistory
// @Summary 创建DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "创建DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deployHistory/createDeployHistory [post]
func CreateDeployHistory(c *gin.Context) {
	var deployHistory model.DeployHistory
	_ = c.ShouldBindJSON(&deployHistory)
	_,total := service.GetDeployHistoryInfoTotal(deployHistory.AppId)
	if total == 0 {
		deployHistory.Status= "install"
	} else {
		deployHistory.Status= "upgrade"
	}
	deployInfo := GetDeployInfo(deployHistory.AppId)
	_,appInfo := service.GetAppsByAppId(deployHistory.AppId)
	deployHistory.DeployInfo = deployInfo
	fileName := fmt.Sprintf("/go/src/gin-vue-admin/k8s_file/tmp/deploy-%v.yaml",deployHistory.AppId)
	dstFile,err := os.Create(fileName)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(deployInfo)
	result := doDeploy(deployHistory.Status, appInfo.AppName, appInfo.Namespace ,fileName)
	deployHistory.Result = result
	err = service.CreateDeployHistory(deployHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func doDeploy(status string,appName string,namespace string,fileName string) (result string) {
	if status == "install"{
		cmd := fmt.Sprintf("cd /go/src/gin-vue-admin/k8s_file/gwx-app-with-istio-ingress&& helm %v -n %v  %v . -f %v --debug  --dry-run",status,namespace, appName, fileName)
		fmt.Printf("当前命令行: %v\n", cmd)
		_, out, err := gocommand.NewCommand().Exec(cmd)
		if err != nil {
			log.Panic(err)
		}
		if out  == "" {
			log.Printf("\n应用: %v 预执行失败\n执行结果: %v\n",appName,out)
			result = "fail"
		}else {
			cmd = fmt.Sprintf("cd /go/src/gin-vue-admin/k8s_file/gwx-app-with-istio-ingress&& helm %v -n %v  %v . -f %v --debug ",status,namespace, appName, fileName)
			_, out, _ := gocommand.NewCommand().Exec(cmd)
			if out  == "" {
				log.Printf("\n应用: %v 执行失败\n执行结果: %v\n",appName,out)
				result = "fail"
			}else {
				log.Printf("\n应用: %v 执行成功\n执行结果: %v\n",appName,out)
				result = "success"
			}
		}
	}else if status == "upgrade"{
		cmd := fmt.Sprintf("cd /go/src/gin-vue-admin/k8s_file/gwx-app-with-istio-ingress&& helm %v -n %v  %v . -f %v --debug  --dry-run",status,namespace, appName, fileName)
		fmt.Printf("当前命令行: %v\n", cmd)
		_, out, err := gocommand.NewCommand().Exec(cmd)
		if err != nil {
			log.Panic(err)
		}
		if out  == "" {
			log.Printf("\n应用: %v 预执行失败\n执行结果: %v\n",appName,out)
			result = "fail"
		}else {
			cmd = fmt.Sprintf("cd /go/src/gin-vue-admin/k8s_file/gwx-app-with-istio-ingress&& helm %v -n %v  %v . -f %v --debug ",status,namespace, appName, fileName)
			_, out, _ := gocommand.NewCommand().Exec(cmd)
			if out  == "" {
				log.Printf("\n应用: %v 执行失败\n执行结果: %v\n",appName,out)
				result = "fail"
			}else {
				log.Printf("\n应用: %v 执行成功\n执行结果: %v\n",appName,out)
				result = "success"
			}
		}
	}else {
		fmt.Printf("错误的执行类型")
	}
	return
}
// @Tags DeployHistory
// @Summary 删除DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "删除DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deployHistory/deleteDeployHistory [delete]
func DeleteDeployHistory(c *gin.Context) {
	var deployHistory model.DeployHistory
	_ = c.ShouldBindJSON(&deployHistory)
	err := service.DeleteDeployHistory(deployHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DeployHistory
// @Summary 批量删除DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deployHistory/deleteDeployHistoryByIds [delete]
func DeleteDeployHistoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteDeployHistoryByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DeployHistory
// @Summary 更新DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "更新DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /deployHistory/updateDeployHistory [put]
func UpdateDeployHistory(c *gin.Context) {
	var deployHistory model.DeployHistory
	_ = c.ShouldBindJSON(&deployHistory)
	err := service.UpdateDeployHistory(&deployHistory)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DeployHistory
// @Summary 用id查询DeployHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployHistory true "用id查询DeployHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /deployHistory/findDeployHistory [get]
func FindDeployHistory(c *gin.Context) {
	//GetDeployInfo(1)
	fmt.Printf("请求IP: %v\n", c.Request.Header)
	var deployHistory model.DeployHistory
	_ = c.ShouldBindQuery(&deployHistory)
	err, redeployHistory := service.GetDeployHistory(deployHistory.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"redeployHistory": redeployHistory}, c)
	}
}

// @Tags DeployHistory
// @Summary 分页获取DeployHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeployHistorySearch true "分页获取DeployHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deployHistory/getDeployHistoryList [get]
func GetDeployHistoryList(c *gin.Context) {
	var pageInfo request.DeployHistorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetDeployHistoryInfoList(pageInfo)
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
