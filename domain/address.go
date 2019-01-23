package domain

// 地址表
type Address struct {
	ID              int    `xorm:"'id' not null pk autoincr INTEGER"` //
	UID             int    `xorm:"'uid' not null INTEGER"`            //玩家ID
	NickName        string `xorm:"VARCHAR(512)"`                      //玩家昵称
	Phone           int    `xorm:"BIGINT"`                            //手机号码
	Home            string `xorm:"VARCHAR(512)"`                      //所在地
	DetailedAddress string `xorm:"VARCHAR(512)"`                      //详细地址
	Label           string `xorm:"VARCHAR(512)"`                      //标签
	State           int    `xorm:"SMALLINT"`                          //标签
}

// TableName 用户表名
func (Address) TableName() string {
	return "address"
}
