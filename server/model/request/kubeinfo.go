package request


type PodSearch struct {
	Namespace string `json:"namespace"`
	AppName string `json:"app_name"`
	PodEnv string `json:"pod_env"`
}
