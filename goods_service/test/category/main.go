package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"goods_service/proto"
	"goods_service/test"
)

func main() {
	test.InitRPCConnect()
	//TestGetAllCategoriesList()
	//TestGetSubCategory()
	TestCreateCategory()
}

func TestGetAllCategoriesList() {
	response, err := test.GoodsClient.GetAllCategoriesList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.JsonData)
}

func TestGetSubCategory() {
	response, err := test.GoodsClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id:    130358,
		Level: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestCreateCategory() {
	response, err := test.GoodsClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		Name:           "test",
		ParentCategory: 0,
		Level:          1,
		IsTab:          true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
