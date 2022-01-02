package main

import (
	"fmt"
	"goredistest/goroutine"
	"strconv"
	"testing"
	"time"
	
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 增加一个闭包方法
		//第一个参数是用来到时候取闭包方法返回的数据的
		//第二个参数是闭包，处理主要逻辑的地方
		//第三个参数是需要传入到闭包内部的参数，可以是多个，一般建议是map，这样闭包内部好处理
		goroutine.AddGo("kkk"+strconv.Itoa(i), goroutineFun, i)
		goroutine.AddGo("kkkOk"+strconv.Itoa(i), goroutineFun, strconv.Itoa(i), "ok")
	}
	//这里是等待协程处理结束，并获取闭包内部的值
	waitValue := goroutine.Wait()
	fmt.Println(waitValue)

	//模拟使用闭包返回的数据
	if waitValue["kkk1"] == nil {
		fmt.Println("错误")
	} else {
		arr := waitValue["kkk1"].([]interface{})
		for _, v := range arr {
			str := v.(int) + 100
			fmt.Println("res,test1:", str)
		}
	}
}

func goroutineFun(funcValue ...interface{}) interface{} {
	fmt.Println(funcValue)
	time.Sleep(time.Second)
	return funcValue
}
