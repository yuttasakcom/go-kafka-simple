package adapter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yuttasakcom/go-kafka-simple/src/core/common"
)

type FiberApp struct {
	*fiber.App
}

func NewFiberApp() *FiberApp {
	return &FiberApp{fiber.New(fiber.Config{
		CaseSensitive:         false,
		StrictRouting:         false,
		DisableStartupMessage: true,
		ReadBufferSize:        10240,
		ReadTimeout:           30 * time.Second,
	})}
}

func NewFiberHandler(handler func(common.ContextHanlder)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
		return nil
	}
}

type FiberContext struct {
	*fiber.Ctx
}

func NewFiberContext(c *fiber.Ctx) *FiberContext {
	return &FiberContext{c}
}

func (c *FiberContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberContext) Status(code int) common.ContextHanlder {
	c.Ctx.Status(code)
	return c
}

func (c *FiberContext) JSON(v interface{}) error {
	return c.Ctx.JSON(v)
}

func (r *FiberApp) Post(path string, handler func(common.ContextHanlder)) {
	r.App.Post(path, NewFiberHandler(handler))
}
