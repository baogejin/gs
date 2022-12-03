package main

import (
	"fmt"
	"gs/proto/myproto"
)

func main() {
	req := &myproto.LogoutREQ{}
	data, _ := req.Marshal()
	fmt.Println(data)

}
