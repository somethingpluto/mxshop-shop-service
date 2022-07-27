package initialize

import (
	"Shop_service/user_service/global"
	"Shop_service/user_service/model"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitDB() {
	var err error
	user := global.ServiceConfig.MysqlInfo.User
	password := global.ServiceConfig.MysqlInfo.Password
	name := global.ServiceConfig.MysqlInfo.Name
	host := global.ServiceConfig.MysqlInfo.Host
	port := global.ServiceConfig.MysqlInfo.Port
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	newLogger := logger.New(log.New(dest, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      false,
	})
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	err = global.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	zap.S().Infow("数据库连接成功")
}
