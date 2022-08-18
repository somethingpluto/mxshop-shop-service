package main

import (
	"context"
	"fmt"
	"goods_service/proto"
	"goods_service/test"
)

func main() {
	test.InitRPCConnect()
	//TestGoodsList()
	//TestBatchGetGoods()
	TestGetGoodsDetail()
}

func TestGoodsList() {
	response, err := test.GoodsClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		IsHot:       false,
		Pages:       1,
		PagePerNums: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Total)
	for _, goods := range response.Data {
		fmt.Println(goods.Name, goods.ShopPrice)
	}
}

func TestBatchGetGoods() {
	response, err := test.GoodsClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{Id: []int32{421, 422}})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Total)
	for _, goods := range response.Data {
		fmt.Println(goods.Name, goods.ShopPrice)
	}
}

func TestCreateGoods() {
	test.GoodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		GoodsSn:         "",
		Stocks:          0,
		MarketPrice:     0,
		ShopPrice:       0,
		GoodsBrief:      "",
		GoodsDesc:       "",
		ShipFree:        false,
		Images:          nil,
		DescImages:      nil,
		GoodsFrontImage: "",
		IsNew:           false,
		IsHot:           false,
		OnSale:          false,
		CategoryId:      0,
		BrandId:         0,
	})
}

func TestGetGoodsDetail() {
	response, err := test.GoodsClient.GetGoodsDetail(context.Background(), &proto.GoodsInfoRequest{Id: 421})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
