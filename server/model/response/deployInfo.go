package response

type SwitchInfo struct {
	GatewaySwitch bool
	HpaSwitch bool
	LivenessSwitch bool
	ReadinessSwitch bool
	MetricsSwitch bool
	CommandSwitch bool
	ArgsSwitch bool
}

type HpaMode struct {
	Enabled bool `yaml:"enabled"`
	MinReplicas int `yaml:"minReplicas"`
	MaxReplicas int `yaml:"maxReplicas"`
	CpuTargetAverageUtilization int `yaml:"cpuTargetAverageUtilization"`
	MemTargetAverageValue	string `yaml:"memTargetAverageValue"`
}
type GatewayHosts struct {
	Hosts	string	`yaml:"hosts"`
}
type GatewayData struct {
	Hosts	string	`yaml:"hosts"`
	Paths	string	`yaml:"paths"`
	WeightOnline int	`yaml:"weight_online"`
	WeightCanary int	`yaml:"weight_canary"`
}
type Image struct {
	Repository	string	`yaml:"repository"`
	TagOnline	string	`yaml:"tagOnline"`
	TagCanary	string	`yaml:"tagCanary"`
	PullPolicy	string	`yaml:"pullPolicy"`
	ContainerPort	int	`yaml:"containerPort"`
	WeightOnline	int	`yaml:"weightOnline"`
	WeightCanary	int	`yaml:"weightCanary"`
	ReplicaCountOnline	int	`yaml:"replicaCountOnline"`
	ReplicaCountCanary	int	`yaml:"replicaCountCanary"`
	CommandEnable	bool	`yaml:"commandEnable"`
	CommandInfo []string	`yaml:"commandInfo,flow"`
	ArgsEnable	bool	`yaml:"argsEnable"`
	ArgsInfo	[]string	`yaml:"argsInfo,flow"`
}

type ServiceAccount struct {
	Create	bool	`yaml:"create"`
	Annotations	map[string]string `yaml:"annotations"`
	Name	string `yaml:"name"`
}
type Gateway struct {
	Enabled	bool `yaml:"enabled"`
	//GatewayHosts []GatewayHosts	`yaml:"gateway_hosts,flow"`
}

type HostAlias struct {
	Enabled	bool `yaml:"enabled"`
}
type HostAliasData struct {
	HostName string `yaml:"hostname"`
	Ip	string	`yaml:"ip"`
}
type Metrics struct {
	Enabled	bool `yaml:"enabled"`
	Port	int `yaml:"port"`
	Path	string	`yaml:"path"`
}
type Service struct {
	ServiceType string `yaml:"type"`
	Port	int	`yaml:"port"`
}
type Liveness struct {
	Enabled bool `yaml:"enabled"`
	Path	string	`yaml:"path"`
	FailureThreshold	int	`yaml:"failureThreshold"`
	InitialDelaySeconds	int	`yaml:"initialDelaySeconds"`
	PeriodSeconds	int	`yaml:"periodSeconds"`
	TimeoutSeconds	int	`yaml:"timeoutSeconds"`
	SuccessThreshold	int	`yaml:"successThreshold"`
}

type Readiness struct {
	Enabled	bool `yaml:"enabled"`
	FailureThreshold	int	`yaml:"failureThreshold"`
	InitialDelaySeconds	int	`yaml:"initialDelaySeconds"`
	Path	string	`yaml:"path"`
	PeriodSeconds	int	`yaml:"periodSeconds"`
	SuccessThreshold	int	`yaml:"successThreshold"`
	TimeoutSeconds	int	`yaml:"timeoutSeconds"`
}
type PodHostAlias struct {
	Hostname	string	`yaml:"hostNames"`
	Ip	string	`yaml:"ip"`
}
type Limits struct {
	Cpu	string	`yaml:"cpu"`
	Memory	string	`yaml:"memory"`
}
type Requests struct {
	Cpu	string	`yaml:"cpu"`
	Memory	string	`yaml:"memory"`
}
type Resources struct {
	Limits
	Requests
}
type DeployInfo struct {
	CurrentEnv string `yaml:"currentEnv"`
	Namespace string `yaml:"namespace"`
	AppName string	`yaml:"appname"`
	HpaMode	`yaml:"hpaMode"`
	Image
	ImagePullSecrets	[]string	`yaml:"imagePullSecrets,flow"`
	NameOverride	string	`yaml:"nameOverride"`
	FullnameOverride	string	`yaml:"fullnameOverride"`
	ServiceAccount	`yaml:"serviceaccount"`
	Metrics
	Readiness
	Service
	Liveness
	HostAlias	`yaml:"hostAlias"`
	PodHostAlias []PodHostAlias `yaml:"podhostAlias"`
	Gateway
	GatewayHosts []GatewayHosts `yaml:"gateway_hosts"`
	GatewayData	[]GatewayData `yaml:"gateway_data"`
	Resources
	NodeSelector	map[string]string `yaml:"nodeSelector"`
	Tolerations	[]string	`yaml:"tolerations"`
	PodSecurityContext	map[string]string	`yaml:"podSecurityContext"`
	SecurityContext	map[string]string `yaml:"securityContext"`
}
