package main

import (
	"context"
	"fmt"
	"goods_service/proto"
	"goods_service/test"
	"time"
)

func main() {
	//TestBrandList()
	TestCreateBrand()
}

// TestBrandList
// @Description: 测试商品列表
//
func TestBrandList() {
	test.InitRPCConnect()
	response, err := test.GoodsClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       1,
		PagePerNums: 10,
	})
	if err != nil {
		panic(err)
	}
	for _, brand := range response.Data {
		fmt.Println(brand)
	}
	fmt.Println("品牌总条数", response.Total)
}

func TestCreateBrand() {
	test.InitRPCConnect()
	name := "test" + time.Now().String()
	logo := "logo" + time.Now().String()
	response, err := test.GoodsClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: name,
		Logo: logo,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
