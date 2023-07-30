package router

import (
	"net/http"

	"github.com/yuttasakcom/go-kafka-simple/src/core/adapter"
	"github.com/yuttasakcom/go-kafka-simple/src/core/app"
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
)

func Register(app *app.App) {
	health := adapter.NewFiberHandler(func(ch common.ContextHanlder) {
		ch.Status(http.StatusOK).JSON(map[string]string{"status": "ok"})
	})
	app.Get("/system/health", health)
}
