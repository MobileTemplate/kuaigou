
package api

import (
	"qianuuu.com/lib/util"
	"qianuuu.com/kuaigou/usecase"
	"github.com/labstack/echo/middleware"
	"qianuuu.com/kuaigou/login/internal/rest"
)

var (
	uc        *usecase.Usecase
	dc util.Map
)

// Init 初始化
func Init(usecase *usecase.Usecase) {
	uc = usecase
}

// Destroy 销毁
func Destroy() {
}

// SetupRoutes 设置路由规则
func SetupRoutes(r rest.Router) {
	e := r.GetEcho()

	e.Use(middleware.Recover())
	e.Use(rest.MiddleAddHead())
	e.Use(rest.MiddleLogger())

	e.Post("/users/login", UserLogin)

}
