package redis_client

import (
	"github.com/go-redis/redis"
	"log"
	"github.com/Andrewpqc/muxi-promotion-service/utils"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     utils.GetRedisAddr(),
		Password: utils.GetRedisPassword(), // no password set
		DB:       0,                        // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("connect redis failed!")
	} else {
		log.Println("connect redis successful!")
	}
}

/**
向"xueer-promotion"有序集合中添加新成员，初始分数为1；若成员已存在
则将分数加一
 */


func MyZadd(id string) (error) {

	//ZAdd,返回成功增加的成员的数量;如果是一个新成员，那么返回１
	//如果是一个是一个旧成员，该成员的分数被覆盖，并且该函数返回0,
	//即成功增加的成员数为0
	added,err:=RedisClient.ZAddNX("xueer-promotion", redis.Z{
		Score:  1,
		Member: id,
	}).Result()
	if added == int64(1) { //新添加一个成员,初始分数为1
		return nil
	} else if added==int64(0) {//重复，则增加分数
		_, err := RedisClient.ZIncr("xueer-promotion", redis.Z{
			Score:  1,
			Member: id,
		}).Result()
		return err
	}
	return err
}

/**
获取得分最高的前几名的成员和分数信息
 */
func GetTopWithScore(top int64) (vals []redis.Z, err error) {
	vals, err = RedisClient.ZRevRangeWithScores("xueer-promotion", 0, top-1).Result()

	return
}

/**
获取某一成员当前的排名情况
*/
func GetRankbyID(id string) (int64, error) {
	zrank := RedisClient.ZRevRank("xueer-promotion", id)
	return zrank.Val()+1, zrank.Err()
}


/**
获取某一排名范围内的用户信息
 */
 func GetRangeWithScore(start int64,end int64)([]redis.Z,error){
 	if end==-1{
 		total,_:=RedisClient.ZCard("xueer-promotion").Result()
 		end=total
	}
 	return RedisClient.ZRevRangeWithScores("xueer-promotion",start-1,end-1).Result()
 }