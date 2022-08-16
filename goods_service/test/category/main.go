package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"goods_service/test"
)

func main() {
	test.InitRPCConnect()
	TestGetAllCategoriesList()
}

func TestGetAllCategoriesList() {
	response, err := test.GoodsClient.GetAllCategoriesList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.JsonData)
}
