package app

import (
	"context"
	"fmt"
	"handling-grpc-errors-in-rest-gateway/internal/gateway/services/calculator"
	"log"

	"github.com/gin-gonic/gin"
)

const grpcHost = "localhost"
const grpcPort = 8081

type app struct {
	port int
}

func New(port int) *app {
	return &app{
		port: port,
	}
}

func (a *app) Run(ctx context.Context) error {
	r := gin.Default()

	calculator.RegisterCalculatorRoutes(r, grpcHost, grpcPort)

	log.Print("Starting gateway at port: ", a.port)
	return r.Run(fmt.Sprintf(":%d", a.port))
}
