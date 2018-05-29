package utils

import (
	"testing"
	"os"
	"time"
	"fmt"
)

func TestGetRedisPassword(t *testing.T) {
	if got := GetRedisPassword(); got != "" {
		t.Errorf("got GetRedisPassword() = %v,but need ''", got)
	}

	os.Setenv("REDIS_PASSWORD", "you never konw")
	if got := GetRedisPassword(); got != "you never konw" {
		t.Errorf("got GetRedisPassword() = %v,but need 'you never konw'", got)
	}
	os.Unsetenv("REDIS_PASSWORD")
}

func TestGetRedisAddr(t *testing.T) {
	if got := GetRedisAddr(); got != "localhost:6379" {
		t.Errorf("got GetRedisAddr() = %v,but need 'localhost:6379'", got)
	}

	os.Setenv("REDIS_ADDR", "xx.xx.xx.xx:1234")
	if got := GetRedisAddr(); got != "xx.xx.xx.xx:1234" {
		t.Errorf("got GetRedisAddr() = %v,but need 'xx.xx.xx.xx:1234'", got)
	}
	os.Unsetenv("REDIS_ADDR")
}

func TestGetBasicAuthInfo(t *testing.T) {

	if got := GetBasicAuthInfo(); got["andrewpqc"] != "andrewpqc" {
		t.Errorf("got GetBasicAuthInfo()['andrewpqc']= %v,but need 'andrewpqc'", got["andrewpqc"])
	}
	os.Setenv("BASIC_AUTH_INFO", "AA:aa,BB:bb")
	got := GetBasicAuthInfo()
	if got["AA"] != "aa" {
		t.Errorf("got GetBasicAuthInfo()['AA'] = %v,but need 'aa'", got["AA"])
	}
	if got["BB"] != "bb" {
		t.Errorf("got GetBasicAuthInfo()['BB'] = %v,but need 'bb'", got["BB"])
	}
	if got["CC"] != "" {
		t.Errorf("got GetBasicAuthInfo()['CC'] = %v,but need ''", got["CC"])
	}
	os.Unsetenv("BASIC_AUTH_INFO")
}


func TestGenerateToken(t *testing.T) {
	token:=GenerateToken("1",60)
	fmt.Println(token)
}

func TestCheckToken(t *testing.T) {
	if flag,id,err:=CheckToken("b9urQ4Iu0aIEOA5CRdzhaw==.0KyxoeqZ7ouwY1FbC27ckA==.rwKLZowI48ul/lAmXe136g==");flag!=2||err!=nil{
		//t.Error("CheckToken error")
		fmt.Println(flag,id,err)
	}
	time.Sleep(6)
	//if flag,_,_:=CheckToken("b9urQ4Iu0aIEOA5CRdzhaw==.0KyxoeqZ7ouwY1FbC27ckA==./9wky9v5iA0eblQbfXNxgQ==");flag!=1{
		//t.Error("token error")
	//}
}