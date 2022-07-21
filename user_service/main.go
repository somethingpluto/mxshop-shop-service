package main

import (
	"Shop_service/user_service/handler"
	"Shop_service/user_service/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	Port := flag.Int("port", 3000, "端口号")
	flag.Parse()
	fmt.Println("ip:", *IP)
	fmt.Println("port:", *Port)

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserService{})
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc: " + err.Error())
	}
}
