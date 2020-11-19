package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/response"
)

func GetDeployInfoByAppId(appId uint) (err error, deployInfo response.DeployInfo) {
	err = global.GVA_DB.Table("apps").Select("apps.*,deploy.args_switch,deploy.command_switch,deploy.container_port,deploy.pull_policy,deploy.replica_count_canary,deploy.replica_count_online,deploy.repository,deploy.tag_canary,deploy.tag_online,deploy.weigit_canary,deploy.weigit_online,deploy.weight_canary,deploy.weight_online,hpa.cpu_target_average_utilization,hpa.max_replicas,hpa.mem_target_average_value,hpa.min_replicas,liveness.failure_threshold,liveness.initial_delay_seconds,liveness.path,liveness.period_seconds,liveness.success_threshold,liveness.timeout_seconds,readiness.failure_threshold,readiness.initial_delay_seconds,readiness.path,readiness.period_seconds,readiness.success_threshold,readiness.timeout_seconds,metrics.path,metrics.port,resources.cpu_limit,resources.cpu_requests,resources.mem_limit,resources.mem_requests,namespaces.namespace").Where("apps.id = ?", appId).Joins("left join namespaces on namespaces.id = apps.namespace_id").Joins("left join deploy on deploy.app_id = apps.id").Joins("left join hpa on hpa.app_id = apps.id").Joins("left join liveness on liveness.app_id = apps.id").Joins("left join readiness on readiness.app_id = apps.id").Joins("left join metrics on metrics.app_id = apps.id").Joins("left join resources on resources.app_id = apps.id").Scan(&deployInfo).Error
	if err != nil {
		fmt.Printf("\nerr ......... : %v\n\n ",err)
	}
	return
}

func GetDeploySwitchByAppId(appId uint) (err error, switchInfo response.SwitchInfo) {
	err = global.GVA_DB.Table("apps").Select("apps.gateway_switch,apps.hpa_switch,apps.liveness_switch,apps.readiness_switch,apps.metrics_switch,deploy.command_switch,deploy.args_switch").Where("apps.id = ?", appId).Joins("left join deploy on deploy.app_id = apps.id").Scan(&switchInfo).Error
	if err != nil {
		fmt.Printf("\nerr ......... : %v\n\n ",err)
	}
	return
}
