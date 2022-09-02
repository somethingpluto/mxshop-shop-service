package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"inventory_service/proto"
)

var InventoryClient proto.InventoryClient

func init() {
	var err error
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	InventoryClient = proto.NewInventoryClient(conn)
}

func main() {
	//TestSetInv()
	//TestInvDetail()
	//TestSell()
	TestReback()
}

func TestSetInv() {
	_, err := InventoryClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: 421,
		Num:     100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置库存成功")
}
func TestInvDetail() {
	response, err := InventoryClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: 421,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

func TestSell() {

	_, err := InventoryClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{{GoodsId: 421, Num: 10}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("销售成功")
}

func TestReback() {
	_, err := InventoryClient.ReBack(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{{GoodsId: 421, Num: 10}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("退回成功")
}
