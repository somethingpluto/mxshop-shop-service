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
	//TestGetGoodsDetail()
	//TestCreateGoods()
	//TestDeleteGoods()
	TestUpdateGoods()
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
	response, err := test.GoodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		GoodsSn:         "111",
		Stocks:          1111,
		MarketPrice:     1111,
		ShopPrice:       1111,
		GoodsBrief:      "111",
		GoodsDesc:       "111",
		ShipFree:        false,
		Images:          []string{"11111111"},
		DescImages:      []string{"1111111"},
		GoodsFrontImage: "111111",
		IsNew:           false,
		IsHot:           false,
		OnSale:          false,
		CategoryId:      136851,
		BrandId:         934,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetGoodsDetail() {
	response, err := test.GoodsClient.GetGoodsDetail(context.Background(), &proto.GoodsInfoRequest{Id: 421})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestDeleteGoods() {
	response, err := test.GoodsClient.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{Id: 847})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Success)
}

func TestUpdateGoods() {
	response, err := test.GoodsClient.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:   846,
		Name: "更行好吧",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
