package server

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	slog "github.com/Sellsuki/sellsuki-go-logger"
	"github.com/yuttasakcom/go-kafka-simple/src/core/app"
	"github.com/yuttasakcom/go-kafka-simple/src/core/config"
	"github.com/yuttasakcom/go-kafka-simple/src/core/database"
	"github.com/yuttasakcom/go-kafka-simple/src/core/router"
)

type Serverer interface {
	Start()
}

type Server struct {
	config config.Configer
	store  *database.Store
}

func NewServer(config config.Configer) *Server {
	return &Server{
		config: config,
		store:  database.NewStore(config.DB()),
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
		slog.L().Info("Gracefully shutting down...")
		app.Shutdown()
	}()

	go func() {
		defer wg.Done()
		router.Register(app, s.store)
		host := s.config.App().Host()
		app.Listen(host)
	}()

	wg.Wait()
}
