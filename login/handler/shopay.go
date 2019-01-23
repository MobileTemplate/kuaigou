
// 支付相关接口实现

package handler

import (
	"time"
	"qianuuu.com/lib/values"
	"qianuuu.com/kuaigou/usecase"
)

type ShopayImpl struct {
	Uc        *usecase.Usecase
}

func CreateShopayImps(uc *usecase.Usecase) *ShopayImpl {
	return &ShopayImpl{
		Uc:        uc,
	}
}

func (si *ShopayImpl) CreateOrder(uid int, gvm values.ValueMap, opeid int, payname, paytype, appid string) (values.ValueMap, error) {
	//goodtype string, fee int,
	//	count, gift,
	//fee := gvm.GetInt("fee")
	//count := gvm.GetInt("count")
	//gift := gvm.GetInt("gift")
	//goodtype := gvm.GetString("goodtype")
	//si.Uc.CreateOrder(uid, fee, payname, "client", count, gift, opeid, appid, goodtype, paytype)
	return nil, nil
}
func (si *ShopayImpl) OrderSuccess(orderid int, paytime time.Time, Goodtype string, cashback1, cashback2 float64) error {
	//si.Uc.OrderSuccess(orderid, paytime, Goodtype, 0, 0)
	return nil
}
func (si *ShopayImpl) OrderParems(orderid int, params values.ValueMap) error {
	//si.Uc.OrderParems(orderid, params)
	return nil
}


func (si *ShopayImpl) GetGoodConfig(uid int, gid string) (values.ValueMap, error) {
	// 测试金额订单
	//if gid == "999" {
	//	ret := values.ValueMap{
	//		"fee":      100,
	//		"count":    1,
	//		"gift":     0,
	//		"goodtype": "diamond",
	//		"id":       999,
	//	}
	//	return ret, nil
	//}
	//products, err := si.Uc.ConfigFile(uid, "products")
	//if err != nil {
	//	return nil, err
	//}
	//confg := products.GetString("value")
	//goods, err := values.NewValueMapArray([]byte(confg))
	//if err != nil {
	//	return nil, err
	//}
	//var good values.ValueMap
	//for _, g := range goods {
	//	if g.GetString("id") == gid {
	//		good = g
	//	}
	//}
	//// 检测配置文件是否包含所有必要的 key
	//if !good.HasKey("gift") ||
	//	!good.HasKey("count") ||
	//	!good.HasKey("goodtype") {
	//	logs.Error("%v", good)
	//	return nil, errors.New("商城缺少配置项")
	//}
	return nil, nil
}

func (si *ShopayImpl) GetOrderByTradeNo(tradeno string) (values.ValueMap, error) {
	//si.Uc.GetOrderByTradeNo(tradeno)
	return nil, nil
}
