
package rest

import (
	"net/http"

	"github.com/labstack/echo"
	"qianuuu.com/lib/values"
)

type handler func(tools Tools) error

// Router 路由功能模块
type Router interface {
	Run(addr string)

	HEAD(string, handler)
	GET(string, handler)
	POST(string, handler)
	OPTIONS(string, handler)

	GetEcho() *echo.Echo
}

// Tools http 请求封装
type Tools interface {
	ParamInt(name string) int
	ParamString(name string) string
	FormInt(name string) int
	FormString(name string) string
	BodyText() ([]byte, error)
	BodyValueMap() (values.ValueMap, error)
	GetIP() string
	OK(it interface{}) error
	BadRequest(format string, a ...interface{}) error
	String(int, string) error

	Request() *http.Request
}

// New 创建
func New() Router {
	e := echo.New()
	return &echoRouterImpl{
		e: e,
	}
}
