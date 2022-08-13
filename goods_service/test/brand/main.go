package main

import (
	"context"
	"fmt"
	"goods_service/proto"
	"goods_service/test"
)

func main() {
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
