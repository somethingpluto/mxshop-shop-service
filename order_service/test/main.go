package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"order_service/proto"
)

var OrderServiceClient proto.OrderClient

func init() {
	var err error
	conn, err := grpc.Dial("192.168.8.1:52472", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	OrderServiceClient = proto.NewOrderClient(conn)
}

func main() {
	//TestCreateCartItem(26, 20, 424) // success
	//TestCartItemList(26) // success
	//TestUpdateCartItem(26, 424, true, 20) // success
	//TestDeleteCartItem(26, 421) // success
	//TestOrderList(1, 1, 10) // success
	//TestCreateOrder() // success
	//TestOrderList(26, 1, 10) // success
	TestOrderDetail() //success
}

func TestCreateCartItem(userId int32, nums int32, goodsId int32) {
	response, err := OrderServiceClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  userId,
		Nums:    nums,
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestCartItemList(userId int32) {
	response, err := OrderServiceClient.CartItemList(context.Background(), &proto.UserInfo{Id: userId})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestUpdateCartItem(userId int32, goodsId int32, checked bool, nums int32) {
	response, err := OrderServiceClient.UpdateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  userId,
		GoodsId: goodsId,
		Checked: true,
		Nums:    nums,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestDeleteCartItem(userId int32, goodsId int32) {
	response, err := OrderServiceClient.DeleteCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  userId,
		GoodsId: goodsId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestOrderList(userId int32, pages int32, pagePerNums int32) {
	response, err := OrderServiceClient.OrderList(context.Background(), &proto.OrderFilterRequest{
		UserId:      userId,
		Pages:       pages,
		PagePerNums: pagePerNums,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestCreateOrder() {
	response, err := OrderServiceClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  26,
		Address: "??????????????????",
		Name:    "pluto",
		Mobile:  "1234567",
		Post:    "?????????",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestOrderDetail() {
	response, err := OrderServiceClient.OrderDetail(context.Background(), &proto.OrderRequest{
		Id:     1,
		UserId: 26,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
