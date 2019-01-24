package usecase

import (
	"fmt"

	"qianuuu.com/kuaigou/domain"
)

func (uc *Usecase) UserLogin(phone int) (*domain.User, error) {
	user := &domain.User{}
	has, err := uc.client.WhereGet(user, " phone=?", phone)
	if !has {
		return nil, fmt.Errorf("用户不存在！")
	}
	return user, err
}

func (uc *Usecase) UserRegister(phone int, pwd string) (*domain.User, error) {
	ss := uc.client.NewSession()
	defer ss.Close()
	ss.Begin()
	user := &domain.User{
		Phone:    phone,
		Password: pwd,
		Sex:      1,
		Icon:     "https://webzy.oss-cn-hangzhou.aliyuncs.com/kg_shop/icon/kg_icon.png",
	}
	user.SetMisc(user.GetMisc())
	id, err := uc.client.GetNextSerial(user)
	if err != nil {
		ss.Rollback()
		return nil, fmt.Errorf("生成账号失败请确认")
	}
	user.ID = id
	user.NickName = fmt.Sprintf("kg_%d", id)
	retc, err := ss.Insert(user)
	if err != nil {
		ss.Rollback()
		return nil, err
	} else if retc != 1 {
		ss.Rollback()
		return nil, fmt.Errorf("注册失败请确认填写信息是否正确")
	}
	ss.Commit()
	return user, nil
}

func (uc *Usecase) UserInfo(uid int) (*domain.User, error) {
	user := &domain.User{
		ID: uid,
	}
	has, err := uc.client.Get(user)
	if !has {
		return nil, fmt.Errorf("用户不存在！")
	}
	return user, err
}

func (uc *Usecase) UserUpdate(uid int) (*domain.User, error) {
	user := &domain.User{
		ID: uid,
	}
	has, err := uc.client.Get(user)
	if !has {
		return nil, fmt.Errorf("用户不存在！")
	}
	return user, err
}
