package global

import (
	"context"
	"github.com/Deansquirrel/goToolMSSql"
)

const (
	//PreVersion = "1.0.5 Build20190821"
	//TestVersion = "0.0.0 Build20190101"
	Version = "0.0.0 Build20190101"
	Type    = "ServiceSupportHelper"
)

var (
	HeartBeatCron      = "0 * * * * ?"
	RefreshSvrInfoCron = "0 0/5 * * * ?"
)

var Ctx context.Context
var Cancel func()
var DbConfig *goToolMSSql.MSSqlConfig

var (
	ClientId      = ""
	ClientType    = ""
	ClientVersion = ""
	HttpAddress   = ""
	HostName      = ""
	InternetIp    = ""
	DbId          = -1
	DbType        = -1
)
