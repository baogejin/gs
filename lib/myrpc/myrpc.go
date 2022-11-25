package myrpc

import (
	"errors"
	"fmt"
	"gs/lib/myconfig"
	"gs/lib/mylog"
	"gs/lib/myutil"
	rpc_logic "gs/server/logic/rpc"
	"net"
	"sync"
	"time"
)

type MyRpc struct {
	server  *Server
	clients sync.Map
	name    string
	address string
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

func (this *MyRpc) NewRpcServer(name string) int32 {
	this.name = name
	this.server = NewServer()
	wait := make(chan bool, 1)
	port := myconfig.Get().RpcPortStart
	go func() {
		for {
			address := fmt.Sprintf("%s:%d", myutil.GetLocalIP(), port)
			ln, err := net.Listen("tcp", address)
			if err != nil {
				port++
				time.Sleep(10 * time.Millisecond)
				continue
			}
			this.address = address
			wait <- true
			this.server.Serve(ln)
		}
	}()
	<-wait
	mylog.Info("node ", this.name, " rpc server start address:", this.address)
	return port
}

func (this *MyRpc) Register(rcvr interface{}) error {
	if this.server == nil {
		return errors.New("server is nil,need NewRpcServer first")
	}
	return this.server.Register(rcvr)
}

func (this *MyRpc) NewRpcClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		mylog.Error(err)
		return
	}
	cli := NewClient(conn)
	this.clients.Store(addr, cli)
	str := "hello world"
	reply := &rpc_logic.LogicAck{}
	err = cli.Call("RpcLogic.Logic", &rpc_logic.LogicReq{MsgId: 1, Data: []byte(str)}, reply)
	if err == nil {
		mylog.Info(string(reply.Data))
	} else {
		mylog.Error(err)
	}

}
