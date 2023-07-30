package core

import (
	"os"

	"github.com/yuttasakcom/go-kafka-simple/src/core/config"
)

type Core struct {
	Cfg config.Configer
}

func NewCore() *Core {
	cfg := config.NewConfig(envFile())
	return &Core{
		Cfg: cfg,
	}
}

func envFile() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}
