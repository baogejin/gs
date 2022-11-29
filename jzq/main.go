package main

import (
	"fmt"
	"gs/lib/myredis"
)

func main() {
	pwd := myredis.GetInstance().HGet(myredis.Account, "sssss")
	fmt.Println(pwd)
}
