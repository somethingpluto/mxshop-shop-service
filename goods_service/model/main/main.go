package main

import (
	"goods_service/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func main() {
	var err error
	dsn := "root:chx200205173214@tcp(120.25.255.207:3306)/mxshop_goods_service?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建日志文件
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.Category{}, &model.Brand{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})
	if err != nil {
		panic(err)
	}
}
