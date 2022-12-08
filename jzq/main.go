package main

import (
	"fmt"
	"gs/lib/myticker"
	"time"
)

func main() {
	myticker.GetInstance().AddTicker(time.Second, test)
	time.Sleep(time.Second * 5)
	myticker.GetInstance().Destory()
	time.Sleep(time.Hour)
}

func test() {
	fmt.Println("1")
}
