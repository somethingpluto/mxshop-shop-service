package main

import "goods_service/initialize"

func main() {
	// 初始化文件路径
	initialize.InitFileAbsPath()
	// 初始化配置文件
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLogger()
	// 初始化数据库
	initialize.InitDB()
}
