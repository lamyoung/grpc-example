package main

import (
	"context"
	"http/service"
	"log"

	"google.golang.org/grpc"
)

func main() {
	//conn, err := grpc.Dial(":9005", grpc.WithTransportCredentials(util.GetClientCredentials()))
	conn, err := grpc.Dial(":9005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接失败，原因：%v", err)
	}
	defer conn.Close()
	orderClient := service.NewOrderServiceClient(conn)
	orderResponse, err := orderClient.GetOrder(context.Background(), &service.OrderReuqest{OrderId: 123})
	if err != nil {
		log.Fatalf("请求收不到返回：%v", err)
	}
	log.Println(orderResponse.OrderId)
}
