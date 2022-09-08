package mode

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"user_service/handler"
	"user_service/proto"
)

// DebugMode
// @Description: debug开发模式
// @param server
// @param ip
// @param port
//
func DebugMode(server *grpc.Server, ip string, port int) {
	proto.RegisterUserServer(server, &handler.UserService{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		zap.S().Errorw("net.Listen错误", "err", err.Error())
		return
	}
	zap.S().Infof("服务启动成功 端口 %s:%d", ip, port)
	err = server.Serve(listen)
	if err != nil {
		zap.S().Errorw("server.Serve错误", "err", err.Error())
		return
	}

}
