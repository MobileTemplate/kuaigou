
package shopay

import (
	"time"

	"qianuuu.com/lib/values"
	"qianuuu.com/player/domain"
)

type Handler interface {
	CreateOrder(int, values.ValueMap, int, string, string, string) (*domain.Order, error)
	OrderSuccess(orderid int, paytime time.Time, Goodtype string, cashback1, cashback2 float64) error
	OrderParems(orderid int, params values.ValueMap) error
	GetOrderByTradeNo(tradeno string) (*domain.Order, error)
	RefreshUserInfo(uid int) error
	GetGoodConfig(uid int, gid string) (values.ValueMap, error)
	GetUserWxOpenid(uid int) (string, error)
	PayIsOpen(paytype, paychannel string) error
}
