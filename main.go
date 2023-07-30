package main

import (
	"os"

	"github.com/yuttasakcom/go-kafka-simple/src/core/config"
	"github.com/yuttasakcom/go-kafka-simple/src/core/server"
)

func main() {
	// @TODO: init logger
	// @TODO: init tracer
	// @TODO: connected db
	// @TODO: start server
	cfg := config.NewConfig(envFile())
	server.NewServer(cfg).Start()
	// @TODO: start worker
	// @TODO: start gRPC server
}

func envFile() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}
