package domain

// 用户表
type User struct {
	ID       int    `xorm:"'id' not null pk autoincr INTEGER"` //玩家ID
	NickName string `xorm:"VARCHAR(512)"`                      //玩家昵称
	Password string `xorm:"VARCHAR(512)"`                      //账号密码
	Icon     string `xorm:"VARCHAR(2048)"`                     //玩家头像
	Sex      int    `xorm:"SMALLINT"`                          //玩家性别
	Phone    int    `xorm:"BIGINT"`                            //手机号码
	Vip      int    `xorm:"SMALLINT"`                          //vip等级
	Money    int    `xorm:"BIGINT"`                            //账号余额
	Misc     string `xorm:"JSONB"`                             //扩展字段
}

// TableName 用户表名
func (User) TableName() string {
	return "users"
}

type UserMisc struct {
	Code string `json:"code"`
}
