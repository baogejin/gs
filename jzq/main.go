package main

import (
	"fmt"
	"gs/define"
	"gs/lib/eventbus"
)

func main() {
	eventbus.GetInstance().Subscribe(define.EventTest, testfun)
	eventbus.GetInstance().Publish(define.EventTest)
	eventbus.GetInstance().Unsubscribe(define.EventTest, testfun)
	eventbus.GetInstance().Publish(define.EventTest)
}

func testfun() {
	fmt.Println("111111")
}
