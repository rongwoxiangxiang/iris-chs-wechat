package main

import (
	"chs/config"
	"chs/lib/service"
	"github.com/pelletier/go-toml"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	config.InitConfig()
	sources := config.Conf.Get("source")
	if sources == nil {
		log.Println("Init application orm failed: database source null")
		return
	}
	config.InitStoreDb(sources.(*toml.Tree))

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化grpc Server
	s := grpc.NewServer()
	// 注册HelloService
	service.RegisterActivityServiceServer(s, service.ActivityServiceImpl{})
	log.Println("Listen on " + Address)
	s.Serve(listen)
}
