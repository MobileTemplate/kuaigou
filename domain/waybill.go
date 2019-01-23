package domain

// 运单表
type Waybill struct {
	ID            int    `xorm:"'id' not null pk autoincr INTEGER"` //运单ID
	OrderNumber   string `xorm:"VARCHAR(512)"`                      //订单号
	UID           int    `xorm:"'uid' not null INTEGER"`            //玩家ID
	GID           int    `xorm:"'gid' not null INTEGER"`            //商品ID
	PayMoney      int    `xorm:"INTEGER"`                           //支付金额
	GoodNumber    int    `xorm:"INTEGER"`                           //商品数量
	WaybillNumber string `xorm:"VARCHAR(512)"`                      //运单号
	BuyersInfo    string `xorm:"JSONB"`                             //收家信息
	GoodInfo      string `xorm:"JSONB"`                             //商品信息
	Misc          string `xorm:"JSONB"`                             //扩展字段
	State         int    `xorm:"SMALLINT"`                          //运单状态

}

// TableName 用户表名
func (Waybill) TableName() string {
	return "waybill"
}
