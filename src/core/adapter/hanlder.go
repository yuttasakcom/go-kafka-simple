package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
)

func NewHandler(handler func(common.ContextHanlder)) fiber.Handler {
	return NewFiberHandler(handler)
}
