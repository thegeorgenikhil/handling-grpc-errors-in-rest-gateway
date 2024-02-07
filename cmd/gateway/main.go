package main

import (
	"context"
	"handling-grpc-errors-in-rest-gateway/internal/gateway/app"
	"log"
)

const (
	port = 8082
)

func main() {
	app := app.New(port)
	if err := app.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
