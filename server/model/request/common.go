package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type CreateApp struct {
	ArgsInfo           string `json:"argsInfo"`
	ArgsSwitch         *bool  `json:"argsSwitch"`
	CommandInfo        string `json:"commandInfo" form:"commandInfo"`
	CommandSwitch      *bool  `json:"commandSwitch"`
	ContainerPort      string    `json:"containerPort"`
	PullPolicy         string `json:"pullPolicy"`
	Repository         string `json:"repository"`
	CpuLimit    string `json:"cpuLimit"`
	CpuRequests string `json:"cpuRequests"`
	MemLimit    string `json:"memLimit"`
	MemRequests string `json:"memRequests"`
}