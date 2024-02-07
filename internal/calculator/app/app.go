package app

import (
	"context"
	"fmt"
	"handling-grpc-errors-in-rest-gateway/internal/calculator/server"
	"handling-grpc-errors-in-rest-gateway/internal/calculator/service"
	pb "handling-grpc-errors-in-rest-gateway/proto"
	"net"

	"log"

	"google.golang.org/grpc"
)

const grpcPort = 8081

type app struct {
}

func New(port int) *app {
	return &app{}
}

func (a *app) Run(ctx context.Context) error {
	// gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Println("not able to listen: ", err)
		return err
	}
	svc := service.NewCourseService()

	s := server.NewCalculatorServer(svc)

	grpc := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(grpc, s)

	log.Println("starting gRPC Server at port: ", grpcPort)

	return grpc.Serve(lis)
}
