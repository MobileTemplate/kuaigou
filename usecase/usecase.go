
// 数据库操作用例
// 提供了所有的数据操作函数，可独立测试

package usecase

import (
	"math/rand"
	"time"

	"qianuuu.com/lib/client"
)

// Usecase 用例
type Usecase struct {
	client *client.Client
}

// NewUsecase 创建一个用例
func NewUsecase(client *client.Client) *Usecase {
	return &Usecase{client: client}
}

// ShowSQL 是否显示 SQL
func (uc *Usecase) ShowSQL(show bool) {
	uc.client.ShowSQL(show)
}

// RandStr 生成随机字符串
func randStr(strlen int) string {
	rand.Seed(time.Now().Unix())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = rand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = rand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}