package calculator

import (
	"fmt"
	pb "handling-grpc-errors-in-rest-gateway/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type calculatorServiceClient struct {
	Client pb.CalculatorServiceClient
}

func NewCalculatorServiceClient(host string, port int) (*calculatorServiceClient, error) {
	calculatorServiceURL := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(calculatorServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return &calculatorServiceClient{
		Client: pb.NewCalculatorServiceClient(conn),
	}, nil
}
