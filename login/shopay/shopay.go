package shopay

import (
	"github.com/labstack/echo"
	"qianuuu.com/kuaigou/login/shopay/paylib"
	"qianuuu.com/lib/logs"
)

var (
	hd Handler
)

type payImpl struct {
}

func (pi *payImpl) Test(c echo.Context) error {
	logs.Info("pay impl ...")
	return nil
}

func Routes(e *echo.Echo, hdv Handler) {
	hd = hdv

	pi := payImpl{}
	e.Get("/shopay/test", pi.Test)

	// 创建
	// createPay("ftafk", "wxh5")

	if hd == nil {
		logs.Error("[shopay] 配置不正确")
		return
	}

	if err := paylib.ParseToml("./shopay.toml"); err != nil {
		logs.Error("[shopay] 配置不正确 %v", err)
		return
	}
	e.Get("/shopay/close/:paytype/users/:id/:gid", UserPayClose)
	e.Get("/shopay/:payname/:paytype/users/:id/:gid", UserPay)
	e.Get("/shopay/:payname/:paytype/users/:id/:gid/H5", UserPay)
	e.Get("/shopay/:payname/notify", NotifyPay)
	e.Post("/shopay/:payname/notify", NotifyPay)
	e.Get("/shopay/:payname/check/:tradeno", CheckPay)
	e.Get("/shopay/check/:tradeno", CheckPay)
	e.Get("/shopay/backurl", BackUrl)
	e.Post("/shopay/backurl", BackUrl)
}
