package rpc_logic

import (
	"gs/lib/mylog"
	"gs/lib/myrpc"
	logic_handler "gs/server/logic/handler"
)

type RpcLogic int

type LogicReq struct {
	myrpc.RpcBaseReq
	Uid   uint64
	MsgId uint32
	Data  []byte
}

type LogicAck struct {
	MsgId uint32
	Data  []byte
}

func (this *RpcLogic) Logic(arg *LogicReq, reply *LogicAck) (err error) {
	msgId, msg := logic_handler.ProcessMsg(arg.Uid, arg.MsgId, arg.Data, arg.NotifyAddr)
	reply.MsgId = uint32(msgId)
	if msg != nil {
		if data, err := msg.Marshal(); err == nil {
			reply.Data = data
		} else {
			mylog.Error(err)
		}
	}
	//test
	// option := xtcp.NewOpts(&myrpc.RpcHandler{}, &myrpc.RpcProtocol{})
	// option.SendBufListLen = 4096
	// client := xtcp.NewConn(option)
	// go client.DialAndServe(arg.Addr)
	// client.SendPacket(&myrpc.RpcPacket{Data: []byte("hello rpc")})
	return
}
