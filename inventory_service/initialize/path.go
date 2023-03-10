package initialize

import (
	"fmt"
	"inventory_service/config"
	"inventory_service/global"
	"path"
	"runtime"
)

func InitFileAbsPath() {
	basePath := getCurrentAbsplutePath()
	global.FilePath = &config.FilePathConfig{
		ConfigFile: basePath + "/config-debug.yaml",
		LogFile:    basePath + "/log",
	}
	fmt.Println("文件路径初始化成功", basePath)
}

func getCurrentAbsplutePath() string {
	var abPath string
	_, fileName, _, ok := runtime.Caller(2)
	if ok {
		abPath = path.Dir(fileName)
	}
	return abPath
}
