
package echotools

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Token JWT
type Token struct {
	token *jwt.Token
	mi    map[string]interface{}
}

// NewToken 创建一个 token
func NewToken(value map[string]interface{}) *Token {
	btk := jwt.New(jwt.SigningMethodHS256)
	btk.Header["typ"] = "JWT"

	return &Token{
		token: btk,
		mi:    value,
	}
}

// SignedString 获取签名
func (t *Token) SignedString(key string, exp time.Duration) (string, error) {
	t.token.Claims["exp"] = time.Now().Add(exp).Unix()
	// 设置其它 claims
	for k, v := range t.mi {
		t.token.Claims[k] = v
	}
	return t.token.SignedString([]byte(key))
}

// createToken 根据内容创建 Token
func createToken(key, content string) *Token {
	t, err := jwt.Parse(content, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err == nil && t.Valid {
		return &Token{
			token: t,
		}
	}
	return nil
}
