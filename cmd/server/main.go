package main

import (
	"fmt"
	"github.com/Michael-Levitin/PingOmetr/internal/delivery"
	"github.com/Michael-Levitin/PingOmetr/internal/logic"
	pb "github.com/Michael-Levitin/PingOmetr/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	host = "localhost"
	port = "5000"
)

func main() {
	// готовимся принимать сообщения от клиента на порту 5000
	var err error
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("server: error starting tcp listener: ", err)
		os.Exit(1)
	}
	log.Println("server: tcp listener started at port: ", port)

	grpcServer := grpc.NewServer() // создаем новый grpc сервер
	pingLogic, err := logic.NewPingLogic()
	if err != nil {
		log.Println("server: failed creating logic:", err)
		os.Exit(1)
	}

	pingServer := delivery.NewPingServer(*pingLogic)   // ... а логику в библиотеку
	pb.RegisterPingOmetrServer(grpcServer, pingServer) // регистрируем сервис библиотеки в grpc

	if err = grpcServer.Serve(lis); err != nil { // передаем полученные от клиента данные
		log.Println("server: error serving grpc: ", err)
		os.Exit(1)
	}
	log.Println("server is running")
}
