package main

import (
	"gs/lib/mylog"
	"gs/lib/myredis"
)

func main() {
	myredis.GetInstance().HMSet("testhash", "1", "2", "2", "3")
	mylog.Info(myredis.GetInstance().HGet("testhash", "1"))
	mylog.Info(myredis.GetInstance().HSetNX("testhash", "3", "2"))
	mylog.Info(myredis.GetInstance().HGetAll("testhash"))
}
