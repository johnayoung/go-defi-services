package main

import (
	"coingecko/handler"
	pb "coingecko/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("coingecko"),
	)

	// Register handler
	pb.RegisterCoingeckoHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
