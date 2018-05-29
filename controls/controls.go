package controls

import (
	"github.com/kataras/iris"
	"xueer-promotion-service/utils"
	"strings"
	"strconv"
	"time"
	"xueer-promotion-service/redis-client"
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
	expries:=ctx.URLParam("ex")
	if !strings.HasPrefix(url,"http"){
		url="https://"+url
	}
	sign := utils.GetSecretKey()
	token := utils.NewJWToken(sign)
	var tokenString string
	if expries!=""{
		print("here is have expries")
		tokenString, _ = token.GenJWToken(map[string]interface{}{
			"id": string(id),
			"current_time":time.Now().Second(),
			"ex":  string(expries),
		})
	}else{
		print("here is don't have expries")
		tokenString, _ = token.GenJWToken(map[string]interface{}{
			"id": string(id),
		})
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"token":"http://127.0.0.1:8080/promotion/?t="+tokenString+"&landing="+url})
}


//处理推广请求api
//啊啊啊啊啊啊啊
//类型断言啊啊啊啊啊
//go语言的类型转换好TM搞人
//浪费这么长时间


func ProcessPromotionRequest(ctx iris.Context){
	usertoken:=ctx.URLParam("t")
	landing:=ctx.URLParam("landing")
	sign := utils.GetSecretKey()
	token := utils.NewJWToken(sign)
	got,_:=token.ParseJWToken(usertoken)
	//fmt.Println(got["id"],got["expires"])

	//先来检查token有没有过期

	////////////////////////////////////////////////////////////////////
	//expries:=got["expires"].(string) 注意此处如果是空字符串，类型断言会报错//
	////////////////////////////////////////////////////////////////////


	ex:=got["ex"] //若没有expires参数，则拿到的是nil
	if ex!=nil{
		ex_int,_:=strconv.Atoi(ex.(string))
		thatTime:=got["current_time"].(float64)
		result := float64(time.Now().Second()) - thatTime
		if result > float64(ex_int) || result < 0 {
			//token过期了
			ctx.StatusCode(iris.StatusForbidden)
			ctx.WriteString("<h1>该链接已经过期了，请联系链接发布者重新获取链接!<\\h1>")
			return
		}
	}

	//向数据库添加该请求的记录
	ID_str,_:=got["id"].(string)
	redis_client.MyZadd(ID_str)
	ctx.Redirect(landing)
}


//获取前x名的用户以及分数
func GetTopX(ctx iris.Context){
	n, _:= ctx.Params().GetInt("n")
	result,_:=redis_client.GetTopWithScore(int64(n))
	ctx.JSON(iris.Map{"data":result})
}


func GetRank(ctx iris.Context){
	id:= ctx.Params().Get("id")
	fmt.Println(id)
	fmt.Println(id=="1")
	a,_:=redis_client.GetRankbyID(id)
	ctx.JSON(iris.Map{"rank":a,"id":id})
}

func CleanDB(ctx iris.Context){
	data,err:=redis_client.RedisClient.Del("xueer-promotion").Result()
	ctx.JSON(iris.Map{"data":data,"err":err})
	ctx.StatusCode(iris.StatusOK)
}

func GetStatistic(ctx iris.Context){
	data,err:=redis_client.RedisClient.ZCard("xueer-promotion").Result()
	ctx.JSON(iris.Map{"data":data,"err":err})
	ctx.StatusCode(iris.StatusOK)
}