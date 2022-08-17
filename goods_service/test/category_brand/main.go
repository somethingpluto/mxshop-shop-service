package main

import (
	"context"
	"fmt"
	"goods_service/proto"
	"goods_service/test"
)

func main() {
	test.InitRPCConnect()
	//TestCategoryBrandList()
	//TestGetCategoryBrandList()
	//TestCreateCategoryBrand()
	//TestDeleteCategoryBrand()
	TestUpdateCategoryBrand()
}

func TestCategoryBrandList() {
	response, err := test.GoodsClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{
		Pages:       1,
		PagePerNums: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("总数", response.Total)
	fmt.Println(response.Data)
}

func TestGetCategoryBrandList() {
	response, err := test.GoodsClient.GetCategoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: 136846,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("总数", response.Total)
	fmt.Println(response.Data)
}

func TestCreateCategoryBrand() {
	response, err := test.GoodsClient.CreateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		CategoryId: 136846,
		BrandId:    700,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Id)
}

func TestDeleteCategoryBrand() {
	response, err := test.GoodsClient.DeleteCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id: 25799,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Success)
}

func TestUpdateCategoryBrand() {
	response, err := test.GoodsClient.UpdateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id:         25804,
		CategoryId: 136846,
		BrandId:    614,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Success)
}
