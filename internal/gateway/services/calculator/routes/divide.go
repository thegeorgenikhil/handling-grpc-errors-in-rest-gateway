package routes

import (
	grpcErrors "handling-grpc-errors-in-rest-gateway/pkg/grpc_errors"
	pb "handling-grpc-errors-in-rest-gateway/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type DivideRequest struct {
	NumberOne float32 `json:"number_one"`
	NumberTwo float32 `json:"number_two"`
}

func Divide(ctx *gin.Context, cl pb.CalculatorServiceClient) {
	var req DivideRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": true, "message": "invalid request body"})
		return
	}

	res, err := cl.DivideNumbers(ctx, &pb.DivideNumbersRequest{
		NumberOne: req.NumberOne,
		NumberTwo: req.NumberTwo,
	})

	if err != nil {
		s := status.Convert(err)
		ctx.JSON(grpcErrors.GetHttpStatusFromGRPCErrCode(s.Code()), gin.H{"error": true, "message": s.Message()})
		return
	}

	ctx.JSON(int(res.Code), res)
}
