package server

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/yuttasakcom/go-kafka-simple/src/core"
	"github.com/yuttasakcom/go-kafka-simple/src/core/app"
	"github.com/yuttasakcom/go-kafka-simple/src/core/router"
)

type Serverer interface {
	Start()
}

type Server struct {
	core *core.Core
}

func NewServer() *Server {
	return &Server{
		core: core.NewCore(),
	}
}

func (s *Server) Start() {
	app := app.NewApp()

	wg := new(sync.WaitGroup)
	wg.Add(2)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer wg.Done()
		<-c
		fmt.Println("Gracefully shutting down...")
		app.Shutdown()
	}()

	go func() {
		defer wg.Done()
		router.Register(app)
		host := s.core.Cfg.App().Host()
		app.Listen(host)
	}()

	wg.Wait()
}
