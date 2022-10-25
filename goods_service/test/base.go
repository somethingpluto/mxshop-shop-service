package test

import (
	"go.uber.org/zap"
	"goods_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Conn *grpc.ClientConn
var GoodsClient proto.GoodsClient

// InitRPCConnect
// @Description: 初始化GRPC连接
//
func InitRPCConnect() {
	var err error
	Conn, err = grpc.Dial("192.168.8.1:51903", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("测试环境——grpc.Dial失败", "err", err.Error())
		return
	}
	GoodsClient = proto.NewGoodsClient(Conn)
}
