package main

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"user_service/model"
)

var DB *gorm.DB

func main() {
	var err error
	dsn := "root:chx200205173214@tcp(127.0.0.1:3306)/mxshop_user_service?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建日志文件
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      false,
	})

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//err = DB.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic(err)
	//}
	options := &password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: sha512.New,
	}
	salt, encodedPwd := password.Encode("admin", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("pluto%d", i),
			Mobile:   fmt.Sprintf("1517161878%d", i),
			Password: newPassword,
		}
		DB.Save(&user)
	}
}
