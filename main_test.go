package main

import (
	"testing"
	"muxi-promotion-service/redis-client"
	"github.com/kataras/iris/httptest"
	"encoding/json"
	"time"
	"strings"
)

func init() {
	redis_client.RedisClient.Del("xueer_promotion")
}


type Token struct{
	Token string
}


func fromPrivateLink2QueryArgs(tk string)[]string{
	var mm []string
	queryStr:=tk[33:]
	for _,i:=range strings.Split(string(queryStr),"&"){
		a:=strings.Split(i,"=")
		mm=append(mm,a[1])
	}
	return mm
}


func TestBasicAuth(t *testing.T) {
	e := httptest.New(t, newApp())

	//测试不受Basic Auth保护的路由
	e.GET("/").Expect().Status(httptest.StatusOK)
	// without basic
	e.GET("/api/v1.0/statistic/").Expect().Status(httptest.StatusUnauthorized)

	e.GET("/api/v1.0/statistic/").WithBasicAuth("andrewpqc", "andrewpqc").Expect().
		Status(httptest.StatusOK)
	//// with valid basic auth
	//e.GET("/admin").WithBasicAuth("myusername", "mypassword").Expect().
	//	Status(httptest.StatusOK).Body().Equal("/admin myusername:mypassword")
	//e.GET("/admin/profile").WithBasicAuth("myusername", "mypassword").Expect().
	//	Status(httptest.StatusOK).Body().Equal("/admin/profile myusername:mypassword")
	//e.GET("/admin/settings").WithBasicAuth("myusername", "mypassword").Expect().
	//	Status(httptest.StatusOK).Body().Equal("/admin/settings myusername:mypassword")
	//
	//// with invalid basic auth
	//e.GET("/admin/settings").WithBasicAuth("invalidusername", "invalidpassword").
	//	Expect().Status(httptest.StatusUnauthorized)

	redis_client.RedisClient.Del("xueer_promotion")
}

func TestGetPrivatePromotionLin(t *testing.T) {
	e := httptest.New(t, newApp())
	e.GET("/api/v1.0/private-promotion-link/").
		WithQuery("id", "1").WithQuery("url", "www.baidu.com").
		WithBasicAuth("andrewpqc", "andrewpqc").Expect().Status(httptest.StatusOK)
	e.GET("/api/v1.0/private-promotion-link/").
		WithQuery("id", "1").WithQuery("url", "www.baidu.com").WithQuery("ex", "1000").
		WithBasicAuth("andrewpqc", "andrewpqc").Expect().Status(httptest.StatusOK)
	e.GET("/api/v1.0/private-promotion-link/").
		WithQuery("id", "1").WithQuery("ex", "1000").
		WithBasicAuth("andrewpqc", "andrewpqc").Expect().Status(httptest.StatusBadRequest)
	redis_client.RedisClient.Del("xueer_promotion")
}


func TestProcessRequest(t *testing.T) {
	e := httptest.New(t, newApp())
	a:=e.GET("/api/v1.0/private-promotion-link/").
		WithQuery("id", "1").WithQuery("url", "www.baidu.com").
		WithBasicAuth("andrewpqc", "andrewpqc").Expect().Status(httptest.StatusOK).Body().Raw()

	var tk1 Token
	json.Unmarshal([]byte(a),&tk1)
	aa:=fromPrivateLink2QueryArgs(tk1.Token)
	e.GET("/promotion/").WithQuery("t",aa[0]).WithQuery("landing",aa[1]).
		Expect().Status(httptest.StatusPermanentRedirect)



	b:=e.GET("/api/v1.0/private-promotion-link/").
		WithQuery("id", "1").WithQuery("url", "www.baidu.com").WithQuery("ex","5").
		WithBasicAuth("andrewpqc", "andrewpqc").Expect().Status(httptest.StatusOK).Body().Raw()

	var tk2 Token
	json.Unmarshal([]byte(b),&tk2)
	aa=fromPrivateLink2QueryArgs(tk2.Token)
	e.GET("/promotion/").WithQuery("t",aa[0]).WithQuery("landing",aa[1]).
		Expect().Status(httptest.StatusPermanentRedirect)
	time.Sleep(6*time.Second)
	e.GET("/promotion/").WithQuery("t",aa[0]).WithQuery("landing",aa[1]).
		Expect().Status(httptest.StatusForbidden)

	redis_client.RedisClient.Del("xueer_promotion")
}

func TestOtherUnimportant(t *testing.T){
	e := httptest.New(t, newApp())
	e.GET("/api/v1.0/statistic/").WithBasicAuth("andrewpqc","andrewpqc").
		Expect().Status(httptest.StatusOK)
	e.GET("/api/v1.0/clean/").WithBasicAuth("andrewpqc","andrewpqc").
		Expect().Status(httptest.StatusOK)
	redis_client.MyZadd("A")
	redis_client.MyZadd("1")
	redis_client.MyZadd("B")
	redis_client.MyZadd("C")
	redis_client.MyZadd("A")
	redis_client.MyZadd("D")
	redis_client.MyZadd("2")
	redis_client.MyZadd("A")
	e.GET("/api/v1.0/top/{n}/").WithPath("n",2).
		WithBasicAuth("andrewpqc","andrewpqc").Expect().Status(httptest.StatusOK)

	e.GET("/api/v1.0/rank/{id}").WithPath("id","1").
		WithBasicAuth("andrewpqc","andrewpqc").Expect().Status(httptest.StatusOK)

	e.GET("/api/v1.0/range/").WithQuery("start",1).WithQuery("end","3").
		WithBasicAuth("andrewpqc","andrewpqc").Expect().Status(httptest.StatusOK)

	redis_client.RedisClient.Del("xueer_promotion")
}