package main

import (
	"grpc-example/service"
	"grpc-example/util"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	// gwmux := runtime.NewServeMux()
	// opt := []grpc.DialOption{grpc.WithInsecure()}
	// err := service.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:9003", opt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rpcServer := grpc.NewServer(grpc.Creds(util.GetServerCredentials()))
	//rpcServer := grpc.NewServer()
	service.RegisterOrderServiceServer(rpcServer, new(service.OrderService))
	//lis, _ := net.Listen("tcp", ":9005")
	http.ListenAndServeTLS(
		":9005",
		"../cert/server.pem",
		"../cert/server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("收到请求%v", r)
			rpcServer.ServeHTTP(w, r)
		}),
	)
	// http.ListenAndServe(
	// 	":9005",
	// 	grpcHandlerFunc(rpcServer),
	// )
}

// func grpcHandlerFunc(grpcServer *grpc.Server) http.Handler {
// 	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		grpcServer.ServeHTTP(w, r)
// 	}), &http2.Server{})
// }

// func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
// 	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
// 			grpcServer.ServeHTTP(w, r)
// 		} else {
// 			otherHandler.ServeHTTP(w, r)
// 		}
// 	}), &http2.Server{})
// }
