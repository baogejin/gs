package main

import (
	"fmt"
	"gs/lib/myredis"
)

func main() {
	ok := myredis.GetInstance().HSetNX(myredis.Account, "jzq", "234")
	fmt.Println(ok)
}
