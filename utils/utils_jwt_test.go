package utils

import (
	"testing"
)

func TestJWTTool(t *testing.T) {
	// 签名字符串
	sign := GetSecretKey()

	token := NewJWToken(sign)

	// -----------  生成jwt token -----------
	tokenString, _ := token.GenJWToken(map[string]interface{}{
		"name": "root",
		"key":  "1",
	})

	// -----------  解析 jwt token -----------
	got, _ := token.ParseJWToken(tokenString)
	want := "root"
	if got["name"] != want {
		t.Errorf("jwt tool failed!")
	}
	if got["key"] != "1" {
		t.Error("jwt tool error")
	}
}
