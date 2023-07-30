package config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type Configer interface {
	App() *App
}
type Config struct {
	app *App
}

func NewConfig(path string) *Config {
	env, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		app: &App{
			host: env["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(env["APP_PORT"])
				if err != nil {
					log.Fatalf("Error converting APP_PORT to int: %v", err)
				}
				return p
			}(),
		},
	}
}

func (c *Config) App() *App {
	return c.app
}
