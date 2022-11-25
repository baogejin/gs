package myrpc

import "sync"

type MyRpc struct {
	server  *Server
	clients sync.Map
}

var myRpc *MyRpc
var once sync.Once

func GetInstance() *MyRpc {
	once.Do(func() {
		myRpc = new(MyRpc)
		myRpc.init()
	})
	return myRpc
}

func (this *MyRpc) init() {

}
