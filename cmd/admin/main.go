package main

import (
	"context"
	"fmt"
	pb "github.com/Michael-Levitin/PingOmetr/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var (
	host = "localhost"
	port = "5000"
)

func main() {
	// подключаемся к grpc серверу
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println("could not connect to grpc server: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	grpcClient := pb.NewPingOmetrClient(conn) // передаем подключение в сервис

	adminDialer(grpcClient)
}

func adminDialer(grpcClient pb.PingOmetrClient) {
	for {
		data := &pb.GetAdminDataResponse{}
		var err error
		data, err = grpcClient.GetAdminData(context.TODO(), &pb.GetAdminDataRequest{})
		log.Println(data, err)
	}
}
