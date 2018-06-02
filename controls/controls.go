package controls

import (
	"github.com/kataras/iris"
	"github.com/Andrewpqc/muxi-promotion-service/utils"
	"strings"
	"strconv"
	"time"
	"github.com/Andrewpqc/muxi-promotion-service/redis-client"
)

//fmt.Println(strings.HasPrefix("my string", "prefix"))  // false
//fmt.Println(strings.HasPrefix("my string", "my"))      // true

//获取专属推广连接api
func GetPrivatePromotionLink(ctx iris.Context) {
	//get the url params
	id := ctx.URLParam("id")
	url := ctx.URLParam("url")
	if id == "" || url == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	expries := ctx.URLParam("ex")
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	sign := utils.GetSecretKey()
	token := utils.NewJWToken(sign)
	var tokenString string
	if expries != "" {
		tokenString, _ = token.GenJWToken(map[string]interface{}{
			"id":           string(id),
			"current_time": time.Now().Second(),
			"ex":           string(expries),
			"landing":url,
		})
	} else {
		tokenString, _ = token.GenJWToken(map[string]interface{}{
			"id": string(id),
			"landing":url,
		})
	}
	ctx.StatusCode(iris.StatusOK)
	long_url:="https://promotion.andrewpqc.xyz/promotion/?t=" + tokenString
	short_url:=utils.Long2Short(long_url)
	ctx.JSON(iris.Map{"private_promotion_link":short_url})
}

//处理推广请求api
//啊啊啊啊啊啊啊
//类型断言啊啊啊啊啊
//go语言的类型转换好TM搞人
//浪费这么长时间

////////////////////////////////////////////////////////////////////
//expries:=got["expires"].(string) 注意此处如果是空字符串，类型断言会报错//
////////////////////////////////////////////////////////////////////

func ProcessPromotionRequest(ctx iris.Context) {
	usertoken := ctx.URLParam("t")

	sign := utils.GetSecretKey()
	token := utils.NewJWToken(sign)
	got, _ := token.ParseJWToken(usertoken)

	//先来检查token有没有过期
	//若没有expires参数，则拿到的是ni
	ex := got["ex"]
	if ex != nil {
		ex_int, _ := strconv.Atoi(ex.(string))
		thatTime := got["current_time"].(float64)
		result := float64(time.Now().Second()) - thatTime
		if result > float64(ex_int) || result < 0 {
			ctx.StatusCode(iris.StatusForbidden)
			ctx.WriteString("<h1>该链接已经过期了，请联系链接发布者重新获取链接!")
			return
		}
	}

	//向数据库添加该请求的记录
	ID_str, _ := got["id"].(string)
	redis_client.MyZadd(ID_str)
	landing:=got["landing"].(string)
	ctx.StatusCode(iris.StatusPermanentRedirect)
	ctx.Redirect(landing)
}

//获取前x名的用户以及分数
func GetTopX(ctx iris.Context) {
	n, _ := ctx.Params().GetInt("n")
	result, _ := redis_client.GetTopWithScore(int64(n))
	ctx.JSON(iris.Map{"data": result})
}

//获取某一id用户的排名
func GetRank(ctx iris.Context) {
	id := ctx.Params().Get("id")
	a, _ := redis_client.GetRankbyID(id)
	ctx.JSON(iris.Map{"rank": a, "id": id})
}

//清空数据库
func CleanDB(ctx iris.Context) {
	data, _:= redis_client.RedisClient.Del("xueer-promotion").Result()
	ctx.JSON(iris.Map{"data": data})
	ctx.StatusCode(iris.StatusOK)
}

//获取参与此次活动的总人数
func GetStatistic(ctx iris.Context) {
	data, err := redis_client.RedisClient.ZCard("xueer-promotion").Result()
	ctx.JSON(iris.Map{"data": data, "err": err})
	ctx.StatusCode(iris.StatusOK)
}

//范围获取
func GetPageNationInfo(ctx iris.Context) {
	start:=ctx.URLParam("start")
	start_int,_:=strconv.Atoi(start)
	end:=ctx.URLParam("end")
	end_int,_:=strconv.Atoi(end)
	data,_:=redis_client.GetRangeWithScore(int64(start_int),int64(end_int))
	ctx.JSON(iris.Map{"data":data})
	ctx.StatusCode(iris.StatusOK)
}
