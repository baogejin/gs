package main

import (
	"gs/define"
	"gs/lib/eventbus"
	"gs/lib/mylog"
)

func main() {
	eventbus.GetInstance().SubscribeWithTag(define.EventTest, testfun, "1")
	eventbus.GetInstance().SubscribeWithTag(define.EventTest, testfun, "2")
	eventbus.GetInstance().Publish(define.EventTest)
	eventbus.GetInstance().UnsubscribeByTag(define.EventTest, "1")
	eventbus.GetInstance().Publish(define.EventTest)
}

func testfun() {
	mylog.Warning("111111")
}
