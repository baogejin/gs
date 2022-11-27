package rpc_logic

import (
	"gs/lib/mylog"
	"gs/lib/myrpc"

	"github.com/xfxdev/xtcp"
)

type RpcLogic int

type LogicReq struct {
	myrpc.RpcBaseReq
	MsgId uint32
	Data  []byte
}

type LogicAck struct {
	MsgId uint32
	Data  []byte
}

func (this *RpcLogic) Logic(arg *LogicReq, reply *LogicAck) (err error) {
	reply.MsgId = arg.MsgId
	reply.Data = arg.Data
	mylog.Infof("%s", arg.Data)

	//test
	option := xtcp.NewOpts(&myrpc.RpcHandler{}, &myrpc.RpcProtocol{})
	option.SendBufListLen = 4096
	client := xtcp.NewConn(option)
	go client.DialAndServe(arg.Addr)
	client.SendPacket(&myrpc.RpcPacket{Data: "hello rpc"})

	return
}
