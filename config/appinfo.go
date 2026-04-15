package config

import (
	"runtime"
	"time"
)

type AppInfo struct {
	AppName   string
	Version   string
	OS        string
	ARCH      string
	Date      string
	RunTime   string
	GoVersion string
}

var appInfo AppInfo

func GetAppInfo() *AppInfo {
	return &appInfo
}

func InitAppInfo(version, date string) {
	// Using a custom app name for my personal fork
	appInfo.AppName = "Lucky-Personal"
	appInfo.Version = version
	appInfo.Date = date
	appInfo.OS = runtime.GOOS
	appInfo.ARCH = runtime.GOARCH
	// Record the startup time in a human-readable local format
	appInfo.RunTime = time.Now().Local().Format("2006-01-02 15:04:05")
	appInfo.GoVersion = runtime.Version()

	buildTime, err := time.Parse("2006-01-02T15:04:05Z", date)
	if err != nil {
		return
	}
	appInfo.Date = buildTime.Local().Format("2006-01-02 15:04:05")
}
