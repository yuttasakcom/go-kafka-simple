package main

import (
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
	"github.com/yuttasakcom/go-kafka-simple/src/core/config"
	"github.com/yuttasakcom/go-kafka-simple/src/core/server"
)

func main() {
	// @TODO: init logger
	// @TODO: init tracer
	// @TODO: connected db
	// @TODO: start server
	cfg := config.NewConfig(common.EnvFile())
	server.NewServer(cfg).Start()
	// @TODO: start worker
	// @TODO: start gRPC server
}
