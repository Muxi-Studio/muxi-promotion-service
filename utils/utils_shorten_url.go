package utils

import (
	"net/http"
	"strings"
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
)

type Resp struct{
	CreateAt string `json:"createAt"`
	Id string `json:"id"`
	Target string `json:"target"`
	Password bool `json:"password"`
	ShortUrl string `json:"shortUrl"`
	Reuse bool `json:"reuse"`
}


func Long2Short(long string)string{
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://kutt.it/api/url/submit", strings.NewReader("target="+long))
	if err != nil {
		// handle error
	}
	X_API_KEY:=os.Getenv("X_API_KEY")
	if X_API_KEY==""{
		log.Fatal("'X_API_KEY' is not in your path, please set it first!")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-API-KEY", X_API_KEY)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var r Resp
	json.Unmarshal(body,&r)
	return r.ShortUrl
}