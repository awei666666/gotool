package main

import (
	"fmt"
	"goredistest/goredis"
	"testing"
)

func TestRedisSet(t *testing.T) {
	key := "key.0101"
	b,e := goredis.SetString(key, "8989")
	fmt.Println("res:", b,e)

	v,e := goredis.GetString(key)
	fmt.Println("get:", v,e)

	_ = goredis.SetExString(key, "0202", 60)

	_ = goredis.SAdd("sadd", "kkk2")

	sAdd, _ := goredis.SMembers("sadd")
	fmt.Println(sAdd)

}


func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
}