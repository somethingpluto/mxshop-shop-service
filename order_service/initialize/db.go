package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"order_service/global"
	"order_service/model"
	"time"
)

func InitDB() {
	var err error
	MySQL := global.ServiceConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", MySQL.User, MySQL.Password, MySQL.Host, MySQL.Port, MySQL.Name)
	// 创建日志文件
	newLogger := logger.New(log.New(logFileWriter, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      false,
	})
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	err = global.DB.AutoMigrate(&model.Inventory{})
	if err != nil {
		zap.S().Errorw("global.DB.AutoMigrate", "err", err.Error())
		panic(err)
	}
	zap.S().Infof("数据库连接成功")
}
