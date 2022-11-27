package main

import (
	"gs/define"
	"gs/lib/mylog"
	"gs/lib/myrpc"
	rpc_logic "gs/server/logic/rpc"
	"time"
)

func main() {
	myrpc.GetInstance().RegisterClient(define.NodeLogic, nil, nil)

	// for {
	ret, err := myrpc.GetInstance().Call(&myrpc.RpcParam{
		Node:   define.NodeLogic,
		Module: "RpcLogic",
		Fn:     "Logic",
		Req:    &rpc_logic.LogicReq{MsgId: 1, Data: []byte("hello world")},
		Ack:    &rpc_logic.LogicAck{},
	})
	if err != nil {
		mylog.Error(err)
	} else {
		ack := ret.(*rpc_logic.LogicAck)
		mylog.Info(string(ack.Data))
	}
	time.Sleep(time.Second * 3)
	// }

	// option := xtcp.NewOpts(&myrpc.RpcHandler{}, &myrpc.RpcProtocol{})
	// option.SendBufListLen = 4096
	// client := xtcp.NewConn(option)
	// go client.DialAndServe("192.168.0.110:13001")
	// _, err := client.SendPacket(&myrpc.RpcPacket{Data: "hello rpc"})
	// mylog.Info(err)
	// time.Sleep(time.Second * 1000)
}
