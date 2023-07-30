package main

import (
	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
	"github.com/yuttasakcom/go-kafka-simple/src/core/config"
	"github.com/yuttasakcom/go-kafka-simple/src/core/server"
)

func main() {
	cfg := config.NewConfig(common.EnvFile())
	initLogger(cfg.App())

	// @TODO: init tracer
	// @TODO: connected db
	// @TODO: start server

	server.NewServer(cfg).Start()
	// @TODO: start worker
	// @TODO: start gRPC server
}

func initLogger(cfg config.App) {
	var level slog.LogLevel = slog.LevelInfo
	if cfg.DebugLog {
		level = slog.LevelDebug
	}

	config := slog.NewProductionConfig()
	config.LogLevel = level
	config.AppName = cfg.AppName
	config.Version = cfg.AppVersion

	slog.L().Configure(config)
}
