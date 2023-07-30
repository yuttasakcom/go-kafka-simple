package main

import "github.com/yuttasakcom/go-kafka-simple/src/core/server"

func main() {
	// @TODO: init logger
	// @TODO: init tracer
	// @TODO: init db
	// @TODO: start server
	server.NewServer().Start()
	// @TODO: start worker
	// @TODO: start gRPC server
}
