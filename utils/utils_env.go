package utils

import (
	"os"
	"strings"
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
func GetSecretKey() string {
	//注意在生产环境需要设置SECRETKEY的值
	if secretKey := os.Getenv("SECRETKEY"); secretKey != "" {
		return secretKey
	} else {
		return "fDEtrkpbQbocVxYRLZrnkrXDWJzRZMfO"//必须32个字符
	}
}

