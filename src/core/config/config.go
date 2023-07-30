package config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type Configer interface {
	App() App
	DB() DB
}
type config struct {
	app App
	db  DB
}

func NewConfig(path string) config {
	env, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return config{
		app: App{
			host: env["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(env["APP_PORT"])
				if err != nil {
					log.Fatalf("Error converting APP_PORT to int: %v", err)
				}
				return p
			}(),
		},
		db: DB{
			Pg: pgDB{
				host: env["PG_DB_HOST"],
				port: func() int {
					p, err := strconv.Atoi(env["PG_DB_PORT"])
					if err != nil {
						log.Fatalf("Error port fail %v", err)
					}
					return p
				}(),
				user:     env["PG_DB_USER"],
				password: env["PG_DB_PASSWORD"],
				dbname:   env["PG_DB_NAME"],
				sslmode:  env["PG_DB_SSLMODE"],
				timezone: env["PG_DB_TIMEZONE"],
			},
			Mg: mgDB{
				host: env["MG_DB_HOST"],
				port: func() int {
					p, err := strconv.Atoi(env["MG_DB_PORT"])
					if err != nil {
						log.Fatalf("Error port fail %v", err)
					}
					return p
				}(),
				user:     env["MG_DB_USER"],
				password: env["MG_DB_PASSWORD"],
				dbname:   env["MG_DB_NAME"],
			},
		},
	}
}

func (c config) App() App {
	return c.app
}

func (c config) DB() DB {
	return c.db
}
