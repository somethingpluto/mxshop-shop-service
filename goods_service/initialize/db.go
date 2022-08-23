package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"goods_service/global"
	"goods_service/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// InitDB
// @Description: 初始化DB
//
func InitDB() {
	var err error
	MySqlInfo := global.ServiceConfig.MySqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", MySqlInfo.User, MySqlInfo.Password, MySqlInfo.Host, MySqlInfo.Port, MySqlInfo.Name)
	// 创建日志文件
	newLogger := logger.New(log.New(logFileWriter, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      false,
	})

	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	err = global.DB.AutoMigrate(&model.Category{}, &model.Brand{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})
	if err != nil {
		panic(err)
	}
	zap.S().Infow("数据库初始化成功")
}
