package controls

import (
	"github.com/kataras/iris"
	"xueer-promotion-service/utils"
	"strings"
	"fmt"
)

//fmt.Println(strings.HasPrefix("my string", "prefix"))  // false
//fmt.Println(strings.HasPrefix("my string", "my"))      // true

//获取专属推广连接api
func GetPrivatePromotionLink(ctx iris.Context){
	//get the url params
	id:=ctx.URLParam("id")
	url:=ctx.URLParam("url")
	if id=="" || url==""{
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	expries:=ctx.URLParam("expries")
	if !strings.HasPrefix(url,"http"){
		url="https://"+url
	}
	sign := utils.GetSecretKey()
	token := utils.NewJWToken(sign)
	var tokenString string
	if expries!=""{
		tokenString, _ = token.GenJWToken(map[string]interface{}{
			"id": string(id),
			"expires":  string(expries),
		})
	}else{
		tokenString, _ = token.GenJWToken(map[string]interface{}{
			"id": string(id),
		})
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"token":"https://xueerpromotion.andrewpqc.xyz/promotion/?t="+tokenString+"&landing="+url})
}


//处理推广请求api
func ProcessPromotionRequest(ctx iris.Context){
	usertoken:=ctx.URLParam("t")
	landing:=ctx.URLParam("landing")
	sign := utils.GetSecretKey()
	token := utils.NewJWToken(sign)
	got,_:=token.ParseJWToken(usertoken)
	fmt.Println(got["id"],got["expries"])
	ctx.Redirect(landing)
}