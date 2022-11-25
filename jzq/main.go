package main

import (
	"gs/lib/myrpc"
)

func main() {
	myrpc.GetInstance().NewRpcClient("192.168.0.110:13000")
}
