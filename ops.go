package anet

const (
	TaskVersionSyncNotReady = iota
	TaskVersionSyncProcessing
	TaskVersionSyncError
	TaskVersionSyncDone
)

type OpsResp struct {
	Time int64  `json:"ts"`
	OK   bool   `json:"ok"`
	Msg  string `json:"msg"`
}

type OpsFullVersionSyncArgs struct {
	Project        string   `json:"project"`          // 项目
	ProjectShort   string   `json:"project_short"`    // 项目段名称
	MaintainTaskID int      `json:"maintain_task_id"` // 任务ID
	SignUrl        string   `json:"sign_url"`         // oss 下载地址
	FileMD5        string   `json:"file_md5"`         // 文件 MD5
	LogPath        string   `json:"log_path"`         // 日志路径
	VerType        int      `json:"ver_type"`         // 版本类型
	Ver            string   `json:"ver"`              // 版本号
	AgentID        string   `json:"agent_id"`         // agent id
	DistList       []string `json:"dist_list"`        // dist 列表
}

type OpsFullVersionSyncStatus struct {
	TaskID         int    `json:"task_id"`          // 任务ID
	TaskMsg        string `json:"task_msg"`         // 任务日志
	TaskDist       string `json:"task_dist"`        // 执行区组
	TaskDistStatus int    `json:"task_dist_status"` // 任务区组执行状态
	TaskDistMsg    string `json:"task_dist_msg"`    // 任务区组执行日志
	Time           string `json:"time"`             // 记录日期
}

type OpsFullVersionSyncRep struct {
	Time int64  `json:"ts"`
	OK   bool   `json:"ok"`
	Msg  string `json:"msg"`
}

type OpsPatchVersionSyncArgs struct {
	Project        string   `json:"project"`          // 项目
	ProjectShort   string   `json:"project_short"`    // 项目短名称
	MaintainTaskID int      `json:"maintain_task_id"` // 任务ID
	SignUrl        string   `json:"sign_url"`         // oss 下载地址
	FileMD5        string   `json:"file_md5"`         // 文件 MD5
	LogPath        string   `json:"log_path"`         // 日志路径
	VerType        int      `json:"ver_type"`         // 版本类型
	Ver            string   `json:"ver"`              // 版本号
	AgentID        string   `json:"agent_id"`         // agent id
	DistList       []string `json:"dist_list"`        // dist 列表
}

type OpsPatchVersionSyncStatus struct {
	TaskID         int    `json:"task_id"`          // 任务ID
	TaskMsg        string `json:"task_msg"`         // 任务日志
	TaskDist       string `json:"task_dist"`        // 执行区组
	TaskDistStatus int    `json:"task_dist_status"` // 任务区组执行状态
	TaskDistMsg    string `json:"task_dist_msg"`    // 任务区组执行日志
	Time           string `json:"time"`             // 记录日期
}

type OpsPatchVersionSyncRep struct {
	Time int64  `json:"ts"`
	OK   bool   `json:"ok"`
	Msg  string `json:"msg"`
}

type OpsVersionUpdateArgs struct {
	Project        string   `json:"project"`          // 项目
	ProjectShort   string   `json:"project_short"`    // 项目短名称
	MaintainTaskID int      `json:"maintain_task_id"` // 任务ID
	VerType        int      `json:"ver_type"`         // 版本类型
	Ver            string   `json:"ver"`              // 版本号
	AgentID        string   `json:"agent_id"`         // agent_id
	DistList       []string `json:"dist_list"`        // dist 列表
}

type OpsVersionUpdateStatus struct {
	TaskID         int    `json:"task_id"`          // 任务ID
	TaskMsg        string `json:"task_msg"`         // 任务日志
	TaskDist       string `json:"task_dist"`        // 执行区组
	TaskDistStatus int    `json:"task_dist_status"` // 任务区组执行状态
	TaskDistMsg    string `json:"task_dist_msg"`    // 任务区组执行日志
	Time           string `json:"time"`             // 记录日期
}

type OpsServiceStartArgs struct {
	Project        string   `json:"project"`          // 项目
	ProjectShort   string   `json:"project_short"`    // 项目短名称
	MaintainTaskID int      `json:"maintain_task_id"` // 任务ID
	VerType        int      `json:"ver_type"`         // 版本类型
	Ver            string   `json:"ver"`              // 版本号
	AgentID        string   `json:"agent_id"`         // agent_id
	DistList       []string `json:"dist_list"`        // dist 列表
}

type OpsServiceStartStatus struct {
	TaskID         int    `json:"task_id"`          // 任务ID
	TaskMsg        string `json:"task_msg"`         // 任务日志
	TaskDist       string `json:"task_dist"`        // 执行区组
	TaskDistStatus int    `json:"task_dist_status"` // 任务区组执行状态
	TaskDistMsg    string `json:"task_dist_msg"`    // 任务区组执行日志
	Time           string `json:"time"`             // 记录日期
}

type OpsServiceStopArgs struct {
	Project        string   `json:"project"`          // 项目
	ProjectShort   string   `json:"project_short"`    // 项目短名称
	MaintainTaskID int      `json:"maintain_task_id"` // 任务ID
	VerType        int      `json:"ver_type"`         // 版本类型
	Ver            string   `json:"ver"`              // 版本号
	AgentID        string   `json:"agent_id"`         // agent_id
	DistList       []string `json:"dist_list"`        // dist 列表
}

type OpsServiceStopStatus struct {
	TaskID         int    `json:"task_id"`          // 任务ID
	TaskMsg        string `json:"task_msg"`         // 任务日志
	TaskDist       string `json:"task_dist"`        // 执行区组
	TaskDistStatus int    `json:"task_dist_status"` // 任务区组执行状态
	TaskDistMsg    string `json:"task_dist_msg"`    // 任务区组执行日志
	Time           string `json:"time"`             // 记录日期
}

type OpsCheckGameStatusArgs struct {
	Project      string `json:"project"`       // 项目
	ProjectShort string `json:"project_short"` // 项目短名称
	Ver          string `json:"ver"`           // 版本号
	AgentID      string `json:"agent_id"`      // agent_id
	DistName     string `json:"dist_name"`     // dist
}

type OpsCheckGameStatusLog struct {
	TaskID         int    `json:"task_id"`          // 任务ID
	TaskMsg        string `json:"task_msg"`         // 任务日志
	TaskDist       string `json:"task_dist"`        // 执行区组
	TaskDistStatus int    `json:"task_dist_status"` // 任务区组执行状态
	TaskDistMsg    string `json:"task_dist_msg"`    // 任务区组执行日志
	Time           string `json:"time"`             // 记录日期
}

type DistProcessStatus struct {
	Ok           bool   `json:"ok"`            // 状态是否正常
	Dist         string `json:"dist_name"`     // dist name
	Host         string `json:"host"`          // 主机
	Service      string `json:"service"`       // 服务
	Ver          string `json:"ver"`           // 版本
	ProcessCount int    `json:"process_count"` // 运行进程数
	ErrMsg       string `json:"err_msg"`       // 错误信息详情
	ErrMsgShort  string `json:"err_msg_short"` // 错误信息
}

type OpsCheckGameStatusRep struct {
	Time       int64             `json:"ts"`          // 返回时间戳
	OK         bool              `json:"ok"`          // 请求是否正常
	Msg        string            `json:"msg"`         // 错误信息
	DistStatus DistProcessStatus `json:"dist_status"` // 区组状态
}
