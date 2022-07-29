package util

import (
	"go.uber.org/zap"
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		zap.S().Errorw("net.ResolveTCPAddr", "err", err.Error())
		return 0, err
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		zap.S().Errorw("net.ListenTCP", "err", err.Error())
		return 0, err
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port, nil
}
