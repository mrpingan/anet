package anet

type distGroup struct {
	AgentId  string   `json:"agent_id"`  // agent id
	DistList []string `json:"dist_list"` // dist 列表
}

type OpsFullVersionSyncArgs struct {
	SignUrl   string      `json:"sign_url"`   // oss 下载地址
	LogPath   string      `json:"log_path"`   // 日志路径
	VerType   int         `json:"ver_type"`   // 版本类型
	DistGroup []distGroup `json:"dist_group"` // 区组组
}

type OpsFullVersionSyncRep struct {
	Time int64  `json:"ts"`
	OK   bool   `json:"ok"`
	Msg  string `json:"msg"`
}
