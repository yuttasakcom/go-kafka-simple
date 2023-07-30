package common

type ContextHanlder interface {
	Bind(v interface{}) error
	Status(code int) ContextHanlder
	JSON(v interface{}) error
	SendString(s string) error
}
