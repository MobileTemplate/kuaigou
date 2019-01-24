package shopay

import (
	"time"

	"qianuuu.com/kuaigou/domain"
	"qianuuu.com/lib/values"
)

type Handler interface {
	CreateOrder(int, values.ValueMap, int, string, string, string) (*domain.Orders, error)
	OrderSuccess(orderid int, paytime time.Time, Goodtype string, cashback1, cashback2 float64) error
	OrderParems(orderid int, params values.ValueMap) error
	GetOrderByTradeNo(tradeno string) (*domain.Orders, error)
	RefreshUserInfo(uid int) error
	GetGoodConfig(uid int, gid string) (values.ValueMap, error)
	GetUserWxOpenid(uid int) (string, error)
	PayIsOpen(paytype, paychannel string) error
}
