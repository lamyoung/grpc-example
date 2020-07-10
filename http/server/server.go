package main

import (
	"context"
	"http/service"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func main() {

	// 基于TCP，不带证书
	// rpcServer := grpc.NewServer()
	// service.RegisterOrderServiceServer(rpcServer, new(service.OrderService))
	// lis, _ := net.Listen("tcp", ":9005")
	// rpcServer.Serve(lis)

	// 基于TCP，带证书
	// rpcServer := grpc.NewServer(grpc.Creds(util.GetServerCredentials()))
	// service.RegisterOrderServiceServer(rpcServer, new(service.OrderService))
	// lis, _ := net.Listen("tcp", ":9005")
	// rpcServer.Serve(lis)

	// 基于http，带证书
	// rpcServer := grpc.NewServer(grpc.Creds(util.GetServerCredentials()))
	// service.RegisterOrderServiceServer(rpcServer, new(service.OrderService))
	// http.ListenAndServeTLS(
	// 	":9005",
	// 	"../cert/server.pem",
	// 	"../cert/server.key",
	// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		log.Printf("收到请求%v", r)
	// 		log.Println(r.ProtoMajor)
	// 		rpcServer.ServeHTTP(w, r)
	// 	}),
	// )

	// 基于h2c，不带证书
	// rpcServer := grpc.NewServer()
	// service.RegisterOrderServiceServer(rpcServer, new(service.OrderService))
	// http.ListenAndServe(
	// 	":9005",
	// 	grpcHandlerFunc(rpcServer),
	// )

	// gRPC+gateway
	rpcServer := grpc.NewServer()
	service.RegisterOrderServiceServer(rpcServer, new(service.OrderService))
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := service.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:9005", opt)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(
		":9005",
		grpcHandlerFunc(rpcServer, gwmux),
	)
}

// func grpcHandlerFunc(grpcServer *grpc.Server) http.Handler {
// 	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		grpcServer.ServeHTTP(w, r)
// 	}), &http2.Server{})
// }

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r)
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
