package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
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
	consulConfig := global.ServiceConfig.ConsulInfo

	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.ServiceConfig.GoodsServiceInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalw("连接 【goods_service】商品服务失败", "err", err)
	}
	zap.S().Infof("goods_service服务连接成功")
	global.GoodsServiceClient = proto.NewGoodsClient(goodsConn)
}

func initInventoryService() {
	consulConfig := global.ServiceConfig.ConsulInfo

	inventoryConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulConfig.Host, consulConfig.Port, global.ServiceConfig.InventoryServiceInfo.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalw("连接 【inventory_service】商品服务失败", "err", err)
	}
	zap.S().Infof("inventory_service服务连接成功")
	global.InventoryServiceClient = proto.NewInventoryClient(inventoryConn)
}
