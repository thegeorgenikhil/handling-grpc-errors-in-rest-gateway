package calculator

import (
	"handling-grpc-errors-in-rest-gateway/internal/gateway/services/calculator/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterCalculatorRoutes(r *gin.Engine, grpcHost string, grpcPort int) {
	svc, err := NewCalculatorServiceClient(grpcHost, grpcPort)
	if err != nil {
		log.Fatalln("not able to create course service client: ", err)
	}

	routes := r.Group("/calculator")

	routes.POST("/divide", svc.Divide)
}

func (svc *calculatorServiceClient) Divide(ctx *gin.Context) {
	routes.Divide(ctx, svc.Client)
}
