package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"inventory_service/proto"
	"sync"
)

var InventoryClient proto.InventoryClient

func init() {
	var err error
	conn, err := grpc.Dial("192.168.8.1:52361", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	InventoryClient = proto.NewInventoryClient(conn)
}

func main() {
	//var wg sync.WaitGroup //var wg sync.WaitGroup
	//TestSetInv()
	//TestInvDetail()

	TestReback()

	//wg.Add(250)
	//for i := 0; i < 250; i++ {
	//	go TestSell(&wg)
	//}
	//wg.Wait()
}

func TestSetInv() {
	for i := 421; i <= 422; i++ {
		_, err := InventoryClient.SetInv(context.Background(), &proto.GoodsInvInfo{
			GoodsId: int32(i),
			Num:     800,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("设置库存成功")
	}

}
func TestInvDetail() {
	response, err := InventoryClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: 421,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func TestSell(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := InventoryClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{{GoodsId: 421, Num: 10}},
	})
	if err != nil {
		fmt.Println(err)
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
