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
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("<h1>Hello, Welcome to use Muxi Promotion Service")
	})
	app.Get("/promotion/",controls.ProcessPromotionRequest)

	// to party
	needAuth := app.Party("/api/v1.0", authentication)
	{
		needAuth.Get("/private-promotion-link/", controls.GetPrivatePromotionLink)
		needAuth.Delete("/clean/",controls.CleanDB)
		needAuth.Get("/statistic/",controls.GetStatistic)
		needAuth.Get("/top/{n:int}/",controls.GetTopX)
		needAuth.Get("/rank/{id:string}/",controls.GetRank)
		needAuth.Get("/range/",controls.GetPageNationInfo)
	}

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"),iris.WithoutVersionChecker)
}

