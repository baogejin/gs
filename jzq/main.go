package main

import (
	"fmt"
	"gs/lib/eventbus"
)

func main() {
	eventbus.GetInstance().Subscribe("1", testfun)
	eventbus.GetInstance().Publish("1")
}

func testfun() {
	fmt.Println("111111")
}
