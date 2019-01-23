
package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/color"
	"qianuuu.com/lib/logs"
	"qianuuu.com/lib/values"
)

// echo 框架的 tools 实现
type echoRouterImpl struct {
	e *echo.Echo
}

func (ei *echoRouterImpl) Run(addr string) {
	// 启用模块
	ei.e.Run(addr)
}

func (ei *echoRouterImpl) HEAD(path string, handler handler) {
	ei.e.Head(path, func(c echo.Context) error {
		return handler(&echoToolsImpl{c: c})
	})
}

func (ei *echoRouterImpl) GET(path string, handler handler) {
	ei.e.Get(path, func(c echo.Context) error {
		return handler(&echoToolsImpl{c: c})
	})
}

func (ei *echoRouterImpl) POST(path string, handler handler) {
	ei.e.Post(path, func(c echo.Context) error {
		return handler(&echoToolsImpl{c: c})
	})
}

func (ei *echoRouterImpl) OPTIONS(path string, handler handler) {
	ei.e.Options(path, func(c echo.Context) error {
		return handler(&echoToolsImpl{c: c})
	})
}

func (ei *echoRouterImpl) GetEcho() *echo.Echo {
	return ei.e
}

type echoToolsImpl struct {
	c       echo.Context
	isbody  bool
	body    []byte
	bodyerr error
}

//ParamInt
func (t *echoToolsImpl) ParamInt(name string) int {
	b, _ := strconv.Atoi(t.c.Param(name))
	return b
}

//ParamString
func (t *echoToolsImpl) ParamString(name string) (value string) {
	return t.c.Param(name)
}

//FormInt
func (t *echoToolsImpl) FormInt(name string) int {
	b, _ := strconv.Atoi(t.c.Form(name))
	return b
}

//FormString
func (t *echoToolsImpl) FormString(name string) string {
	return t.c.Form(name)
}

//BodyText
func (t *echoToolsImpl) BodyText() ([]byte, error) {
	if t.isbody {
		return t.body, t.bodyerr
	}
	t.body, t.bodyerr = ioutil.ReadAll(t.c.Request().Body)
	t.isbody = true
	return t.body, t.bodyerr
}

//BodyValueMap
func (t *echoToolsImpl) BodyValueMap() (values.ValueMap, error) {
	b, err := t.BodyText()
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return nil, errors.New("not body data")
	}
	return values.NewValuesFromJSON([]byte(b))
}

//GetIP
func (t *echoToolsImpl) GetIP() string {
	req := t.c.Request()
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	return remoteAddr
}

// OK 成功
func (t *echoToolsImpl) OK(it interface{}) error {
	b, _ := json.MarshalIndent(it, "", "  ")
	return t.c.String(http.StatusOK, string(b))
}

//BadRequest   错误提示信息
func (t *echoToolsImpl) BadRequest(format string, a ...interface{}) error {
	vp := values.ValueMap{
		"msg": fmt.Sprintf(format, a...),
	}
	t.c.Set("400", vp.GetString("msg"))
	// defer logs.Error("400: %s", vp.GetString("msg"))
	return t.c.String(http.StatusBadRequest, string(vp.ToJSON()))
}

func (t *echoToolsImpl) Request() *http.Request {
	return t.c.Request()
}

func (t *echoToolsImpl) String(code int, content string) error {
	return t.c.String(code, content)
}

// MiddleAddHead 添加响应头
func MiddleAddHead() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Add("Access-Control-Allow-origin", "*")
			c.Response().Header().Add("Content-Type", "application/json")
			c.Response().Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
			c.Response().Header().Add("Access-Control-Max-Age", "3600 * 24")
			c.Response().Header().Add("Access-Control-Allow-Headers", "X-Requested-With, accept, authorization, content-type")
			if err := h(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

// MiddleLogger log 中间件
func MiddleLogger() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			remoteAddr := req.RemoteAddr
			if ip := req.Header.Get("X-Real-IP"); ip != "" {
				remoteAddr = ip
			} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
				remoteAddr = ip
			} else {
				remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
			}

			start := time.Now()
			if err := h(c); err != nil {
				c.Error(err)
			}

			stop := time.Now()
			method := req.Method
			path := req.URL
			size := res.Size()
			n := res.Status()

			// format := "%s %s %s %s %s %d %s"
			// values := []string{}

			code := color.Green(n)
			switch {
			case n >= 500:
				code = color.Red(n)
			case n >= 400:
				code = color.Yellow(n)
			case n >= 300:
				code = color.Cyan(n)
			}

			format := `%s %s %s %s %s %d`
			values := []interface{}{
				remoteAddr, method, code, path, stop.Sub(start), size,
			}
			if n == 400 {
				if v4 := c.Get("400"); v4 != nil {
					format = format + " %s %s"
					values = append(values, "ERR:", c.Get("400"))
				}
			}
			logs.Info(format, values...)

			// logs.Info("%s %s %s %s %s %s", remoteAddr, method, code, path, stop.Sub(start), v4)
			return nil
		}
	}
}
