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
	//TestBannerList()
	//TestCreateBanner()
	//TestDeleteBanner()
	TestUpdateBanner()
}

// TestBannerList
// @Description: 获取轮播图列表
//
func TestBannerList() {
	response, err := test.GoodsClient.BannerList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	for _, banner := range response.Data {
		fmt.Println(banner)
	}
	fmt.Println(response.Total)
}

func TestCreateBanner() {
	response, err := test.GoodsClient.CreateBanner(context.Background(), &proto.BannerRequest{
		Index: 4,
		Image: "http://www.baidu.com",
		Url:   "http://www.baidu.com",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Id)
}

func TestDeleteBanner() {
	response, err := test.GoodsClient.DeleteBanner(context.Background(), &proto.BannerRequest{Id: 5})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestUpdateBanner() {
	response, err := test.GoodsClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    6,
		Index: 5,
		Image: "Test",
		Url:   "Test",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
