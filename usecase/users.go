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
