
package shopay

// PayName 支付通道
type PayName string

const (
	PayNameFtakf  = "ftakf"  // 乾通
	PayNameZfy    = "zfy"    // 掌付云
	PayNameSf     = "sf"     // 顺付
	PayNameAlipay = "alipay" // 阿里支付
)

// PayType 支付类型
type PayType string

const (
	PayTypeWx    = "wxpay" // 微信 SDK 支付
	PayTypeWxH5  = "wxh5"  // 微信 H5 支付
	PayTypeWxgzh = "wxgzh" // 微信公众号支付
	PayTypeAliH5 = "alih5" // 支付宝 H5 支付
	PayTypeQQH5  = "qqh5"  // QQ H5 支付
)
