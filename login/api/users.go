package api

import (
	"github.com/labstack/echo"
	"qianuuu.com/lib/echotools"
)

func UserLogin(c echo.Context) error {
	t := echotools.NewEchoTools(c)
	return t.OK("cg")
}