package response

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type DeploymentInfo struct {
	DeploymentName string `json:"deployment_name"`
	CreationTimestamp v1.Time `json:"creation_timestamp"`
	Replicas int32 `json:"replicas"`
	UnavailableReplicas int32 `json:"unavailable_replicas"`
	PodEnv string `json:"pod_env"`
}

type ContainerInfo struct {
	Name string `json:"name"`
	Ready bool `json:"ready"`
	RestartCount int32 `json:"restart_count"`
}


type PodInfo struct {
	HostIp string `json:"host_ip"`
	NodeName string `json:"node_name"`
	PodIp string `json:"pod_ip"`
	Phase string `json:"phase"`
	ContainerCount int `json:"container_count"`
	Name string `json:"name"`
	CreationTimestamp v1.Time `json:"creation_timestamp"`
	Env string `json:"env"`
	RestartCount int32 `json:"restart_count"`
	ContainersInfo []ContainerInfo
}