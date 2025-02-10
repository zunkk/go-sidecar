package repo

import "strings"

var (
	RootPath = ""

	EnvFilePath = ".env"

	JWTTokenHeaderKey = "token"
)

var (
	AppName = ""

	AppDesc = ""

	Version = ""

	BuildTime = ""

	CommitID = ""

	EnvPrefix = strings.ReplaceAll(strings.ToUpper(AppName), "-", "_")
)

func InitGlobalInfo(appName string, appDesc string, version string, buildTime string, commitID string) {
	AppName = appName
	AppDesc = appDesc
	Version = version
	BuildTime = buildTime
	CommitID = commitID
	EnvPrefix = strings.ReplaceAll(strings.ToUpper(AppName), "-", "_")
}

const (
	cfgFileName = "config.toml"

	debugFileName = "debug-info.json"

	pidFileName = "process.pid"

	logsDirName = "logs"

	IPCFileName = "ipc.sock"
)
