//
// Author: leafsoar
// Date: 2018-06-21 16:30:32
//

// 支付通达管理系统

package paylib

import (
	"fmt"
	"testing"

	"qianuuu.com/lib/logs"
)

func TestPaylib(t *testing.T) {
	fmt.Println("test paylib ...")

	// 读取配置文件
	if err := ParseToml("./shopay.toml"); err != nil {
		logs.Info("配置文件读取失败: %s", err.Error())
		return
	}
	cf := Opts()
	test, err := cf.GetItem("test")
	fmt.Println(test, err)
	fmt.Println(test.CanPayType("qqh5"))
}
