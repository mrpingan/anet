package anet

const (
	SaltInstallDone = iota
	SaltConfigDone
	SaltReboot
)

type SaltInstallArgs struct {
	Master string `json:"master"`
}

type SaltInstallRep struct {
	Time int64  `json:"ts"`
	OK   bool   `json:"ok"`
	Msg  string `json:"msg"`
}
