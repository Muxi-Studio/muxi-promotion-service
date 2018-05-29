package main

import (
	//"time"
	"xueer-promotion-service/utils"
	"xueer-promotion-service/controls"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
)

func newApp() *iris.Application {
	app := iris.Default()

	authConfig := basicauth.Config{
		Users:   utils.GetBasicAuthInfo(),
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		//Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	// to global app.Use(authentication) (or app.UseGlobal before the .Run)
	// to routes
	/*
		app.Get("/mysecret", authentication, h)
	*/

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })
	app.Get("/")
	// to party

	needAuth := app.Party("/api/v1.0", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/private-promotion-link/", controls.GetPrivatePromotionLink)
		// http://localhost:8080/admin/profile

	}

	return app
}

func main() {
	app := newApp()
	// open http://localhost:8080/admin
	app.Run(iris.Addr(":8080"),iris.WithoutVersionChecker)
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	// third parameter it will be always true because the middleware
	// makes sure for that, otherwise this handler will not be executed.

	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}