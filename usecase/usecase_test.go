
package usecase

import (
	"testing"

	"github.com/bmizerany/assert"
	_ "github.com/lib/pq"

	"qianuuu.com/lib/client"
	"qianuuu.com/lib/logs"
)

const UCDBURL = "postgres://postgres:qianuuu_12345@test.qianuuu.cn:5462/jwyl?sslmode=disable"

// 测试用户模块
func TestUsecase(t *testing.T) {
	logs.Info("test player ...")

	client, err := client.NewClient("postgres", UCDBURL)
	assert.Equal(t, err, nil)
	uc := NewUsecase(client)
	defer client.Close()
	_ = uc

}