package utils

import (
	"os"
	"time"
	"encoding/base64"
	"strconv"
	"strings"
	"log"
	"errors"
	"fmt"
)

/**
从环境变量中获取redis的密码
 */
func GetRedisPassword() (string) {
	return os.Getenv("REDIS_PASSWORD") //未设置则为空
}

/**
从环境变量中获取redis服务的地址，若未设置，则默认为"localhost:6379"
 */
func GetRedisAddr() (string) {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		return "localhost:6379"
	} else {
		return addr
	}
}

/**
从环境变量中获取用于Basic auth的用户名和密码，若未设置，则默认为"andrewpqc:andrewpqc"
 */
func GetBasicAuthInfo() (map[string]string) {
	basic_auth_str := os.Getenv("BASIC_AUTH_INFO")
	if basic_auth_str == "" {
		return map[string]string{"andrewpqc": "andrewpqc"}
	} else {
		infoMap := map[string]string{}
		strs := strings.Split(basic_auth_str, ",")
		for _, str := range strs {
			username_password := strings.Split(str, ":")
			infoMap[username_password[0]] = username_password[1]
		}
		return infoMap
	}
}

//注意这里的secretkey一定要是32字节才行
func getSecretKey() string {
	//注意在生产环境需要设置SECRETKEY的值
	if secretKey := os.Getenv("SECRETKEY"); secretKey != "" {
		return secretKey
	} else {
		return "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"
	}
}

//依据传入的id和当前的时间,以及该token过期的时间间隔生成token,
func GenerateToken(id string, expires int64) string {
	key := []byte(getSecretKey())
	goaes := NewGoAES(key)

	userID := []byte(id)
	encrypt1, err := goaes.Encrypt(userID)
	if err != nil {
		log.Fatalln(id, "Generating tokens", err)
	}

	timestamp := []byte(string(time.Now().Unix()))
	encrypt2, err := goaes.Encrypt(timestamp)
	if err != nil {
		log.Fatalln(id, "Generating tokens", err)
	}

	expires_seconds := []byte(string(expires))
	encrypt3, err := goaes.Encrypt(expires_seconds)

	token := base64.StdEncoding.EncodeToString(encrypt1) + "." +
		base64.StdEncoding.EncodeToString(encrypt2) + "." + base64.StdEncoding.EncodeToString(encrypt3)
	return token
}

/**
检查当前传入的token
flag:
	0 token格式错误
	1 token过期
	2 一切正常
 */


func CheckToken(token string) (flag int,id string ,err error) {

	//check token basic format
	tokens := strings.Split(token, ".")
	if len(tokens) != 3 {
		err = errors.New("Invalid token format")
		return 0,"0", err
	}

	// check time
	tokentime, err := base64.StdEncoding.DecodeString(tokens[1])
	if err != nil {
		return 0,"0",err
	}
	print("hhhh",string(tokentime),"hhh")
	fmt.Println(strconv.Atoi(string(tokentime)))

	timestamp, err := strconv.Atoi(string(tokentime))
	if err != nil {
		return 0,"0", err
	}

	//get expires time
	expirestoken,err:=base64.StdEncoding.DecodeString(tokens[2])
	if err!=nil{
		return 0,"0",err
	}
	expires,_:=strconv.Atoi(string(expirestoken))
	print("expires time atoi")
	result := time.Now().Unix() - int64(timestamp)
	if result > int64(expires) || result < 0 {
		err = errors.New("Time Expired")
		return 1,"0",err
	}



	// get id
	uidstr, err := base64.StdEncoding.DecodeString(tokens[0])
	if err != nil {
		return 0,"0",err
	}
	//一切正常，返回用户id
	return 2,string(uidstr),nil
}
