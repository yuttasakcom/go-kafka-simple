package config

import "fmt"

type Apper interface {
	Host() string
}

type App struct {
	host string
	port int
}

func (a App) Host() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}
