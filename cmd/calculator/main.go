package main

import (
	"context"
	"handling-grpc-errors-in-rest-gateway/internal/calculator/app"
	"log"
)

const (
	port = 8080
)

func main() {
	app := app.New(port)
	if err := app.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
