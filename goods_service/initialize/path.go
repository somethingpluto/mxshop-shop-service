package initialize

import (
	"fmt"
	"goods_service/config"
	"goods_service/global"
	"path"
	"runtime"
)

// InitFileAbsPath
// @Description: 初始化文件路径
//
func InitFileAbsPath() {
	basePath := getCurrentAbsolutePath()
	global.FilePath = &config.FilePathConfig{
		ConfigFile: basePath + "/config-debug.yaml",
		LogFile:    basePath + "/log",
	}
	fmt.Println("文件路径初始化成功:", basePath)
}

func getCurrentAbsolutePath() string {
	var abPath string
	_, fileName, _, ok := runtime.Caller(2)
	if ok {
		abPath = path.Dir(fileName)
	}
	return abPath
}
