package utils

import (
	"testing"
)

var longurl="https://www.baidu.com/auth/login/?a=1&b=2&c=balabala"

//环境变量X_API_KEY在CI网站上设置
func TestLong2Short(t *testing.T) {
	short:=Long2Short(longurl)
	if short==""{
		t.Error("Long2Short error")
	}
}
