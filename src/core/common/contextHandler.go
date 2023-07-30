package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type ContextHanlder interface {
	Bind(v interface{}) error
	Status(code int) ContextHanlder
	JSON(v interface{}) error
	SendString(s string) error
	Next() error
	Request() *fasthttp.Request
	Context() *fasthttp.RequestCtx
	Method() string
	Route() *fiber.Route
	OriginalURL() string
	IP() string
	Protocol() string
	Path() string
	Locals(key string, value ...interface{}) (val interface{})
	Response() *fasthttp.Response
}
