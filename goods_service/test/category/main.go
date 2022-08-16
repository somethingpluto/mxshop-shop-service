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
	//TestCreateCategory()
	//TestDeleteCategory()
	TestUpdateCategory()
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
		Name:           "test2",
		ParentCategory: 0,
		Level:          1,
		IsTab:          true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("创建成功", response)
}

func TestDeleteCategory() {
	response, err := test.GoodsClient.DeleteCategory(context.Background(), &proto.DeleteCategoryRequest{Id: 238013})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Success)
}

func TestUpdateCategory() {
	response, err := test.GoodsClient.UpdateCategory(context.Background(), &proto.CategoryInfoRequest{
		Id:             238012,
		Name:           "update",
		ParentCategory: 0,
		Level:          0,
		IsTab:          false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
