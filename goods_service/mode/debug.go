package mode

import (
	"fmt"
	"go.uber.org/zap"
	"goods_service/global"
	"goods_service/handler"
	"goods_service/proto"
	"google.golang.org/grpc"
	"net"
)

func DebugMode(server *grpc.Server, ip string) {
	proto.RegisterGoodsServer(server, &handler.GoodsServer{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, global.FreePort))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}
	zap.S().Infof("服务启动成功 端口 %s:%d", ip, global.FreePort)
	err = server.Serve(listen)
	if err != nil {
		zap.S().Errorw("server.Serve错误", "err", err.Error())
		return
	}

}
