
// echo 帮助

package echotools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"qianuuu.com/lib/logs"
	"qianuuu.com/lib/values"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/color"
)
var isfenghao func(uid int) error
func getFengHao(uid int) (error){
	return isfenghao(uid)
}
func SetFengHao(fn func(uid int) error){
	isfenghao=fn
	//getHeader(1)
}
// EchoTools echo 帮助文件
type EchoTools struct {
	c       echo.Context
	errtext string
	body    string
	isBody  bool
}

// NewEchoTools 帮助
func NewEchoTools(c echo.Context) EchoTools {
	return EchoTools{
		c:      c,
		isBody: false,
	}
}

// ParamInt 获取参数
func (e *EchoTools) ParamInt(name string) int {
	ret, err := strconv.Atoi(e.c.Param(name))
	if e.errtext == "" && err != nil {
		e.errtext = err.Error()
	}
	return ret
}

// ParamString 获取参数
func (e *EchoTools) ParamString(name string) string {
	return e.c.Param(name)
}

// FormInt 获取参数
func (e *EchoTools) FormInt(name string) int {
	ret, err := strconv.Atoi(e.c.Form(name))
	if e.errtext == "" && err != nil {
		e.errtext = err.Error()
	}
	return ret
}
func (e *EchoTools) Method() string {
	return e.c.Request().Method
}
// FormString 获取参数
func (e *EchoTools) FormString(name string) string {
	if !e.isBody {
		txt, _ := e.BodyText()
		e.body = txt
		e.isBody = true
	}
	return strings.TrimSpace(e.c.Form(name))
}

// Errors 获取错误信息
func (e *EchoTools) Errors() error {
	if e.errtext != "" {
		return errors.New(e.errtext)
	}
	return nil
}

// BodyValueMap 从 body 获取数据
func (e *EchoTools) BodyValueMap() (values.ValueMap, error) {
	body, err := e.BodyText()
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, errors.New("not body data")
	}
	return values.NewValuesFromJSON([]byte(body))
}

// BodyText 获取 body 内容
func (e *EchoTools) BodyText() (string, error) {
	if e.isBody {
		return e.body, nil
	}
	e.isBody = true
	body, err := ioutil.ReadAll(e.c.Request().Body)
	if err != nil {
		return "", err
	}
	if len(body) == 0 {
		return "", errors.New("not body data")
	}
	e.body = string(body)
	return e.body, nil
}

func (e *EchoTools) PostFormValue(key string) string {
	return e.c.Request().PostFormValue(key)
}

// NewToken 创建 token
func (e *EchoTools) NewToken(value map[string]interface{}) *Token {
	return NewToken(value)
}

// GetToken 获取 token
func (e *EchoTools) GetToken(key string) (values.ValueMap, error) {
	// 如果 url 中存在 token，则从 url 中提取
	if auth := e.FormString("token"); len(auth) > 0 {
		t, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(key), nil
		})
		if !t.Valid {
			return nil, errors.New("token 验证失败")
		}
		return t.Claims, err
	}
	// url 不存，则从 header 获取
	const Bearer = "Bearer"
	auth := e.c.Request().Header.Get("Authorization")
	l := len(Bearer)

	if len(auth) > l+1 && auth[:l] == Bearer {
		t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(key), nil
		})
		if !t.Valid {
			return nil, errors.New("token 验证失败")
		}
		return t.Claims, err
	}
	return nil, errors.New("获取 token 失败")
}

// VerifyToken 验证传入的
func (e *EchoTools) VerifyToken(key, auth string) (values.ValueMap, error) {
	// 如果 url 中存在 token，则从 url 中提取
	if len(auth) > 0 {
		t, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(key), nil
		})
		if !t.Valid {
			return nil, errors.New("token 验证失败")
		}
		return t.Claims, err
	}

	return nil, errors.New("获取 token 失败")
}

// GetIP 获取 ip 地址
func (e *EchoTools) GetIP() string {
	req := e.c.Request()
	remoteAddr := req.RemoteAddr
	if ipstr := req.Header.Get(echo.XForwardedFor); ipstr != "" {
		ips := strings.Split(ipstr, ", ")
		if len(ips) >= 0 {
			return ips[0]
		}
	}
	if ip := req.Header.Get(echo.XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(echo.XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	return remoteAddr
}

// BadRequest 错误信息
func (e *EchoTools) BadRequest(msg string) error {
	m := map[string]interface{}{
		"msg": msg,
	}
	e.c.Set("400", msg)
	return e.c.JSONIndent(http.StatusBadRequest, m, "", "  ")
}

// OK OK 信息
func (e *EchoTools) OK(state, it interface{}) error {
	ret := values.ValueMap{
		"state": state,
		"data": it,
	}
	return e.c.JSONIndent(http.StatusOK, ret, "", "  ")
}

// MiddleLogger log 中间件
func MiddleLogger(iscolor bool) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			remoteAddr := req.RemoteAddr
			if ip := req.Header.Get(echo.XRealIP); ip != "" {
				remoteAddr = ip
			} else if ip = req.Header.Get(echo.XForwardedFor); ip != "" {
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
			path := req.URL.Path
			if path == "" {
				path = "/"
			}
			size := res.Size()

			n := res.Status()
			code := color.Green(n)
			switch {
			case n >= 500:
				code = color.Red(n)
			case n >= 400:
				code = color.Yellow(n)
			case n >= 300:
				code = color.Cyan(n)
			}
			if !iscolor {
				code = strconv.Itoa(n)
			}

			logs.Info("%s %s %s %s %s %d", remoteAddr, method, path, code, stop.Sub(start), size)
			return nil
		}
	}
}
