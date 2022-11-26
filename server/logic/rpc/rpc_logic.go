package rpc_logic

import (
	"gs/lib/mylog"
)

type RpcLogic int

type LogicReq struct {
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
	return
}
