package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"order_service/global"
	"order_service/proto"
)

func InitOtherService() {
	initGoodsService()
	initInventoryService()
}

func initGoodsService() {
	consulConfig := global.ServiceConfig.Consul

	goodsConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.ServiceConfig.GoodsService),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalw("连接 【goods_service】商品服务失败", "err", err)
	}
	global.GoodsServiceClient = proto.NewGoodsClient(goodsConn)
}

func initInventoryService() {
	consulConfig := global.ServiceConfig.Consul

	inventoryConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.ServiceConfig.InventoryService),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalw("连接 【inventory_service】商品服务失败", "err", err)
	}
	global.InventoryServiceClient = proto.NewInventoryClient(inventoryConn)
}
