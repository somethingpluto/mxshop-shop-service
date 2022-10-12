package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"goods_service/model"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func main() {
	//var err error
	//dsn := "root:chx200205173214@tcp(192.168.8.133:3306)/mxshop_goods_service?charset=utf8mb4&parseTime=True&loc=Local"
	//// 创建日志文件
	//newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	//	SlowThreshold: time.Second,
	//	LogLevel:      logger.Info,
	//	Colorful:      true,
	//})
	//
	//DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//err = DB.AutoMigrate(&model.Category{}, &model.Brand{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})
	//if err != nil {
	//	panic(err)
	//}

	Mysql2Es()
}

func Mysql2Es() {
	var err error
	dsn := "root:chx200205173214@tcp(192.168.8.133:3306)/mxshop_goods_service?charset=utf8mb4&parseTime=True&loc=Local"
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

	host := "http://192.168.8.133:9200"
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	var goods []model.Goods
	DB.Find(&goods)
	for _, g := range goods {
		esModel := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarketPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}

		_, err = client.Index().Index(esModel.GetIndexName()).BodyJson(esModel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
