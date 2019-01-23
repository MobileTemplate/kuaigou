package domain

// 物品表
type Goods struct {
	ID              int    `xorm:"'id' not null pk autoincr INTEGER"` //
	Name            string `xorm:"VARCHAR(512)"`                      //商品昵称
	Type            int    `xorm:"INTEGER"`                           //商品类型
	RawMaterial     string `xorm:"VARCHAR(512)"`                      //原料产地
	Ingredients     string `xorm:"VARCHAR(512)"`                      //配料
	NetContent      string `xorm:"VARCHAR(128)"`                      //净含量
	ShelfLife       int    `xorm:"INTEGER"`                           //保质期
	ProductionNo    string `xorm:"VARCHAR(512)"`                      //生产许可证编号
	Country         string `xorm:"VARCHAR(128)"`                      //国别
	Features        string `xorm:"VARCHAR(256)"`                      //特性
	StorageMethods  string `xorm:"VARCHAR(256)"`                      //存储方法
	Taste           string `xorm:"VARCHAR(128)"`                      //口味
	ProductDetails  string `xorm:"JSONB"`                             //产品详情
	ScrollFigure    string `xorm:"JSONB"`                             //滚动图
	Icon            string `xorm:"VARCHAR(2048)"`                     //icon
	OriginalPrice   int    `xorm:"INTEGER"`                           //原价
	DiscountPrice   int    `xorm:"INTEGER"`                           //折扣价
	RemainingNumber int    `xorm:"INTEGER"`                           //剩余数量
	Misc            string `xorm:"JSONB"`                             //扩展字段
}

// TableName 用户表名
func (Goods) TableName() string {
	return "goods"
}
