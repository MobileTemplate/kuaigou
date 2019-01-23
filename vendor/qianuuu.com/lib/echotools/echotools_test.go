
package echotools

import (
	"fmt"
	"testing"
	"time"

	"qianuuu.com/lib/values"
)

func TestEcho(t *testing.T) {
	// 测试 token
	token := NewToken(func() map[string]interface{} {
		return map[string]interface{}{
			"name": "leafsoar",
		}
	})
	ret, err := token.SignedString("leafsoar", time.Hour*5)
	fmt.Println(ret, err, token)

	vm := values.ValueMap(token.token.Claims)
	fmt.Println(vm.GetString("name"))
}
