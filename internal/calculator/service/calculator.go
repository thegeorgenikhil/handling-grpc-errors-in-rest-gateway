package service

import (
	"context"
	pb "handling-grpc-errors-in-rest-gateway/proto"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CalculatorService struct{}

func NewCourseService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) DivideNumber(ctx context.Context, req *pb.DivideNumbersRequest) (*pb.DivideNumbersResponse, error) {
	if req.NumberTwo == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "division by zero not possible")
	}

	ans := req.NumberOne / req.NumberTwo

	return &pb.DivideNumbersResponse{
		Code: http.StatusOK,
		Answer: ans,
	}, nil
}
