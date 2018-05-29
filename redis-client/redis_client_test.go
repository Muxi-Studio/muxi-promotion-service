package redis_client

import (
	"testing"
)

func TestMyZadd(t *testing.T) {
	if err := MyZadd("A"); err != nil {
		t.Error("MyZadd('A')  call error ")
	}
	if err := MyZadd("B"); err != nil {
		t.Error("MyZadd('B') call error ")
	}
	if err := MyZadd("B"); err != nil {
		t.Error("MyZadd('B') call error ")
	}
	card, _ := RedisClient.ZCard("xueer-promotion").Result()
	if card != int64(2) {
		t.Error("there are not have 2 items in the collection")
	}
	if score, _ := RedisClient.ZScore("xueer-promotion", "A").Result(); score != float64(1) {
		t.Error("error score for itme A")
	}
	if score, _ := RedisClient.ZScore("xueer-promotion", "B").Result(); score != float64(2) {
		t.Error("error score for itme B")
	}
	//清空环境
	RedisClient.Del("xueer-promotion")
}

func TestGetTopWithScore(t *testing.T) {
	MyZadd("A")
	MyZadd("B")
	MyZadd("B")
	MyZadd("B")
	MyZadd("C")
	MyZadd("C")
	MyZadd("D")
	MyZadd("A")
	MyZadd("A")
	MyZadd("A")
	values, _ := GetTopWithScore(3)
	if len(values) != 3 {
		t.Error("the number of the items is incorrect")
	}
	if values[0].Member != "A" || values[1].Member != "B" || values[2].Member != "C" {
		t.Error("the order of the top items is incorrect")
	}
	if values[0].Score != float64(4) || values[1].Score != float64(3) || values[2].Score != float64(2) {
		t.Error("the scores of the top items is incorrect")
	}

	RedisClient.Del("xueer-promotion")
}

func TestGetRankbyID(t *testing.T) {
	MyZadd("A")
	MyZadd("B")
	MyZadd("B")
	MyZadd("B")
	MyZadd("C")
	MyZadd("C")
	MyZadd("D")
	MyZadd("A")
	MyZadd("A")
	MyZadd("A")

	if rank, _ := GetRankbyID("A"); rank != 1 {
		t.Errorf("got GetRankbyID('A') = %v,but want 1", rank)
	}

	if rank, _ := GetRankbyID("D"); rank != 4 {
		t.Errorf("got GetRankbyID('D') = %v,but want 4", rank)
	}
	RedisClient.Del("xueer-promotion")
}
