package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestBasicAuth(t *testing.T) {
	e := httptest.New(t, newApp())

	//测试不受Basic Auth保护的路由
	e.GET("/").Expect().Status(httptest.StatusOK)
	// without basic
	e.GET("/api/v1.0/statistic/").Expect().Status(httptest.StatusUnauthorized)

	e.GET("/api/v1.0/statistic/").WithBasicAuth("andrewpqc","andrewpqc").Expect().
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

}
