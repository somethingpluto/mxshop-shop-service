package initialize

import (
	"Shop_service/user_service/config"
	"Shop_service/user_service/global"
	"fmt"
	"path"
	"runtime"
)

func InitFileAbsPath() {
	basePath := getCurrentAbsolutePath()
	global.FilePath = &config.FilePathConfig{
		ConfigFile: basePath + "/config-debug.yaml",
		LogFile:    basePath + "/log",
	}
	fmt.Println(global.FilePath)
}
func getCurrentAbsolutePath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(2)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
