package domain

// 订单表
type Orders struct {
	ID          int    `xorm:"'id' not null pk autoincr INTEGER"` //订单ID
	UID         int    `xorm:"'uid' not null INTEGER"`            //玩家ID
	GID         int    `xorm:"'gid' not null INTEGER"`            //商品ID
	PayChannel  string `xorm:"VARCHAR(128)"`                      //支付渠道
	PayMoney    int    `xorm:"INTEGER"`                           //支付金额
	GoodsNumber int    `xorm:"INTEGER"`                           //商品数量
	OrderNumber string `xorm:"VARCHAR(512)"`                      //订单号
	Misc        string `xorm:"JSONB"`                             //扩展字段
	State       int    `xorm:"SMALLINT"`                          //订单状态
}

// TableName 用户表名
func (Orders) TableName() string {
	return "orders"
}
