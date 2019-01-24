package api

import (
	"time"

	"fmt"

	"github.com/labstack/echo"
	"qianuuu.com/kuaigou/login/config"
	"qianuuu.com/lib/echotools"
	"qianuuu.com/lib/values"
)

func UserLogin(c echo.Context) error {
	t := echotools.NewEchoTools(c)
	value, err := t.BodyValueMap()
	if err != nil {
		return t.OK(4, "请输入用户名和密码")
	}
	var (
		phone    = value.GetInt("phone")
		code     = value.GetString("code")
		password = value.GetString("password")
	)
	// 通过code登录
	if code != "" {
		token, err := t.VerifyToken(config.Opts.JWTSigning, code)
		if err != nil {
			return t.OK(5, "登录授权已过期")
		}
		if token.GetInt("phone") != phone {
			return t.OK(5, "登录授权已过期")
		}
		phone = token.GetInt("phone")
	}
	user, err := uc.UserLogin(phone)
	if err != nil {
		return t.OK(2, "用户不存在")
	}
	if code == "" {
		//用户存在则确认密码是否正确
		if user.Password != password {
			return t.OK(3, "用户名和密码不正确")
		}
	}

	tmap := map[string]interface{}{
		"id":    user.ID,
		"phone": user.Phone,
	}
	token := t.NewToken(tmap)
	tokenstr, err := token.SignedString(config.Opts.JWTSigning, time.Hour*24)
	ret := values.ValueMap{
		"phone": user.Phone,
		"uid":   user.ID,
		"token": tokenstr,
	}
	return t.OK(1, ret)
}

func UserInfo(c echo.Context) error {
	t := echotools.NewEchoTools(c)
	uid := t.ParamInt("id")
	token, err := t.GetToken(config.Opts.JWTSigning)
	if err != nil {
		return t.OK(5, err.Error())
	}
	if token.GetInt("id") != uid {
		return t.OK(5, "token验证失败")
	}
	fmt.Println("user_info")
	user, err := uc.UserInfo(uid)
	if err != nil {
		return t.OK(2, nil)
	}
	ret := values.ValueMap{
		"phone": user.Phone,
		"uid":   user.ID,
	}
	return t.OK(1, ret)
}
