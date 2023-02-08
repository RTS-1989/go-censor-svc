package main

import (
	"fmt"
	"github.com/RTS-1989/go-censor-svc/pkg/client"
	"github.com/RTS-1989/go-censor-svc/pkg/config"
	"github.com/RTS-1989/go-censor-svc/pkg/db"
	"github.com/RTS-1989/go-censor-svc/pkg/middleware"
	"github.com/RTS-1989/go-censor-svc/pkg/pb"
	"github.com/RTS-1989/go-censor-svc/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	cli := client.InitServiceClient(&c)

	s := services.Server{
		H: h,
		C: cli,
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.LoggingInterceptor))

	pb.RegisterCensorServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
