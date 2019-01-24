//
// Author: leafsoar
// Date: 2018-06-14 09:20:57
//

package paylib

import (
	"time"

	"qianuuu.com/lib/values"
)

// CheckResult 检查返回
type CheckResult struct {
	TradeNo   string          // 商户单号
	IsSucceed bool            // 是否成功
	PayTime   time.Time       // 支付时间
	PayMoney  int             // 支付金额 分
	RetCode   string          // 支付成功内容，用于通知返回
	RetParams values.ValueMap // 支付成功后的返回值
}
