package main

import (
	"context"
	"fmt"
	"goods_service/proto"
	"goods_service/test"
	"time"
)

func main() {
	test.InitRPCConnect()
	//TestBrandList()
	//TestCreateBrand()
	//TestDeleteBrand()
	TestUpdateBrand()
}

// TestBrandList
// @Description: 测试商品列表
//
func TestBrandList() {

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
	name := fmt.Sprintf("test %d", time.Now().Minute())
	logo := fmt.Sprintf("logl %d", time.Now().Minute())
	response, err := test.GoodsClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: name,
		Logo: logo,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestDeleteBrand() {
	response, err := test.GoodsClient.DeleteBrand(context.Background(), &proto.BrandRequest{
		Name: "test 29",
		Logo: "logo 29",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Success)
}

func TestUpdateBrand() {
	response, err := test.GoodsClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   1116,
		Name: "update",
		Logo: "update",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
