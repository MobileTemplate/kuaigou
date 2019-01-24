package shopay

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"qianuuu.com/kuaigou/domain"
	"qianuuu.com/kuaigou/login/shopay/paylib"
	"qianuuu.com/lib/echotools"
	"qianuuu.com/lib/logs"
	"qianuuu.com/lib/values"
)

func createPayHandler(pi *paylib.PayItem) (paylib.PayHandler, error) {
	// 创建通道订单并返回
	var pay paylib.PayHandler

	return nil, errors.New("未知的支付通道")

	return pay, nil
}

// UserPay 用户创建订单
func UserPay(c echo.Context) error {
	t := echotools.NewEchoTools(c)

	// 获取订单参数
	payname := t.ParamString("payname")
	uid := t.ParamInt("id")
	gid := t.ParamString("gid")
	paytype := t.ParamString("paytype")
	debug := t.FormInt("debug")
	ip := t.GetIP()
	if ip == "" {
		ip = "127.0.0.1"
	}
	// 获取通道配置
	pi, err := paylib.Opts().GetItem(payname)
	if err != nil {
		logs.Error("支付文件配置错误: %v", err)
		return t.BadRequest("支付文件配置错误！")
	} else if !pi.IsOpen {
		logs.Error("该支付通道正在维护中: %v", payname)
		return t.BadRequest("该支付通道正在维护，请用其他支付通道！")
	} else if !pi.CanPayType(paytype) {
		logs.Error("该支付方式正在维护: %v", paytype)
		return t.BadRequest("该支付方式正在维护，请用其他支付方式！")
	}
	//判断是否是当前启用的配置
	if debug != 1 {
		temppaytype := paytype[0:1]
		err = hd.PayIsOpen(temppaytype, payname)
		if err != nil {
			logs.Error("不是当前配置通道: %v", err)
			return t.BadRequest("通道未开启！")
		}
	}

	// 获取物品配置
	gvm, err := hd.GetGoodConfig(uid, gid)
	if err != nil {
		logs.Error("商品配置文件错误: %v", err)
		return t.BadRequest("商品配置文件错误！")
	}

	// 创建系统订单
	//fee := gvm.GetInt("fee")
	//gift := gvm.GetInt("gift")
	//count := gvm.GetInt("count")
	//goodtype := gvm.GetString("goodtype")

	logs.Info("[shopay] 创建订单 uid: %d, gid: %v payname: %s, paytype: %s", uid, gid, payname, paytype)
	order, err := hd.CreateOrder(uid, gvm, uid, payname, paytype, pi.AppID)
	if err != nil {
		logs.Error("请求支付失败 uid: %d ：%s", uid, err.Error())
		return t.BadRequest(err.Error())
	}

	// 创建通道订单并返回
	pay, err := createPayHandler(pi)
	if err != nil {
		return t.BadRequest(err.Error())
	}

	// logs.Info("pvm: %v", pvm)
	// logs.Info("pi: %v", pi)
	logs.Info("[shopay] 创建平台订单 uid: %d tradeno: %s paymoney: %v", uid, order.TradeNo, order.PayMoney)

	if paytype == "wxgzh" {
		openid, err := hd.GetUserWxOpenid(uid)
		if err != nil {
			return t.BadRequest(err.Error())
		} else if openid == "" {
			return t.BadRequest("openid为空！")
		}
		parameter := values.ValueMap{
			"openid": openid,
		}
		r, err := pay.CreateUnifiedOrder(paytype, order.TradeNo, order.PayMoney, ip, parameter)
		if err != nil {
			logs.Error("创建平台订单 uid: %d ：%s", uid, err.Error())
			return t.BadRequest(err.Error())
		}
		return t.OK(r)
	}

	// 兼容 返回 web 内容的方式
	params := values.ValueMap{
		"is_content": t.FormString("is_content"),
		"uid":        uid,
		"gid":        gid,
	}
	fmt.Println("params:", params)
	if params.GetString("is_content") == "true" {
		r, err := pay.CreateUnifiedOrder(paytype, order.TradeNo, order.PayMoney, ip, params)
		if err != nil {
			logs.Error("创建平台订单 uid: %d ：%s", uid, err.Error())
			return t.BadRequest(err.Error())
		}
		return c.HTML(http.StatusOK, r.GetString("content"))
	}

	r, err := pay.CreateUnifiedOrder(paytype, order.TradeNo, order.PayMoney, ip, params)
	if err != nil {
		logs.Error("创建平台订单 uid: %d ：%s", uid, err.Error())
		return t.BadRequest(err.Error())
	}
	chkurl := fmt.Sprintf("shopay/%s/check/%s", payname, order.TradeNo)
	ret := values.ValueMap{
		"url":       r.GetString("url"),
		"check_url": chkurl,
		"pay_type":  paytype,
	}
	return t.OK(ret)
}

func CheckPay(c echo.Context) error {
	t := echotools.NewEchoTools(c)

	// 获取订单参数
	payname := t.ParamString("payname")
	tradeno := t.ParamString("tradeno")
	logs.Info("[shopay] check tradeno: %s", tradeno)

	// 获取订单
	order, err := hd.GetOrderByTradeNo(tradeno)
	if err != nil {
		logs.Info("获取订单错误: %s", err.Error())
		return t.BadRequest(err.Error())
	} else if order == nil {
		return t.BadRequest("没有找到此订单 " + tradeno)
	}
	if order.State == 1 {
		return t.OK("该订单已支付成功")
	}
	// 如果没有 payname，则从订单中获取
	if payname == "" {
		payname = order.PayChannel
	}

	// 获取通道配置
	pi, err := paylib.Opts().GetItem(payname)
	if err != nil {
		logs.Error("支付文件配置错误: %v", err)
		return t.BadRequest("支付文件配置错误！")
	}

	// 创建通道订单并返回
	pay, err := createPayHandler(pi)
	if err != nil {
		return t.BadRequest(err.Error())
	}

	result, err := pay.CheckOrder(order)
	if err != nil {
		return t.BadRequest(err.Error())
	}
	if err = actionResult(result, order); err != nil {
		return t.BadRequest(err.Error())
	}
	return t.OK("success")
}

func NotifyPay(c echo.Context) error {
	t := echotools.NewEchoTools(c)
	payname := t.ParamString("payname")
	// 获取通道配置
	pi, err := paylib.Opts().GetItem(payname)
	if err != nil {
		logs.Error("支付文件配置错误: %v", err)
		return t.BadRequest("支付文件配置错误！")
	}

	// 创建通道订单并返回
	pay, err := createPayHandler(pi)
	if err != nil {
		return t.BadRequest(err.Error())
	}

	result, err := pay.NotifyOrder(&t)
	if err != nil {
		return t.BadRequest(err.Error())
	}

	// 获取订单
	order, err := hd.GetOrderByTradeNo(result.TradeNo)
	if err != nil {
		logs.Info("获取订单错误: %s", err.Error())
		return t.BadRequest(err.Error())
	} else if order == nil {
		return t.BadRequest("没有找到此订单 " + result.TradeNo)
	}
	if order.State == 1 {
		//return t.OK("该订单已支付成功")
		return c.String(http.StatusOK, result.RetCode)
	}

	if err = actionResult(result, order); err != nil {
		return t.BadRequest(err.Error())
	}
	return c.String(http.StatusOK, result.RetCode)
}

// 订单处理函数
func actionResult(result *paylib.CheckResult, order *domain.Order) error {
	if result.RetParams != nil {
		err := hd.OrderParems(order.ID, result.RetParams)
		if err != nil {
			logs.Info("添加通道订单信息err:", err)
		}
	}
	// 必须是返回 succeed ，才认为是效验成功，否则都算失败
	if !result.IsSucceed {
		// TODO: 不成功，理应该将订单状态修改 4
		return errors.New("订单检查不成功")
	}
	if result.PayMoney != order.PayMoney {
		return errors.New("订单金额不正确")
	}
	// 史订单成功
	err := hd.OrderSuccess(order.ID, result.PayTime, order.GetMisc().GoodType, 0, 0)
	if err != nil {
		logs.Error("支付检测写入失败： 订单号 :%s %s", order, err.Error())
		return errors.New("支付检测写入数据失败")
	}

	hd.RefreshUserInfo(order.UID)
	return nil
}

// UserPay 用户创建订单
func UserPayClose(c echo.Context) error {
	t := echotools.NewEchoTools(c)
	//ret := values.ValueMap{
	//	"msg": "支付正在维护！",
	//}
	return t.BadRequest("支付正在维护！")
}

// UserPay 用户创建订单
func BackUrl(c echo.Context) error {
	//t := echotools.NewEchoTools(c)
	content := ``
	content += `<script>`
	content += `window.location.href="http://share.kuafuddd.cn/jwyl/index.html";`
	content += `</script>`
	//ret := values.ValueMap{
	//	"msg": "支付正在维护！",
	//}

	//return t.BadRequest("支付正在维护！")
	return c.HTML(http.StatusOK, content)
}
