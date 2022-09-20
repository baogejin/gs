package rpclogic

import (
	"context"
	"fmt"
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

func (this *RpcLogic) Logic(ctx context.Context, arg *LogicReq, reply *LogicAck) (err error) {
	reply.MsgId = arg.MsgId
	reply.Data = arg.Data
	fmt.Printf("%s\n", arg.Data)
	return
}
