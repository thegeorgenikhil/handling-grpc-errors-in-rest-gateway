package server

import (
	"context"
	"handling-grpc-errors-in-rest-gateway/internal/calculator/service"
	pb "handling-grpc-errors-in-rest-gateway/proto"
)

type Server struct {
	service *service.CalculatorService
}

// The following line is for better development experience, so that all the unimplemented methods and their params can be seen in the same file (by hovering) instead of having to go to the app.go file.
var server *Server = nil
var _ pb.CalculatorServiceServer = server

func NewCalculatorServer(calculatorService *service.CalculatorService) *Server {
	return &Server{
		service: calculatorService,
	}
}

func (s *Server) DivideNumbers(ctx context.Context, req *pb.DivideNumbersRequest) (*pb.DivideNumbersResponse, error) {
	return s.service.DivideNumber(ctx, req)
}
