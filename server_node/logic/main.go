package logic

import (
	"fmt"
	"gs/define"
	myrpc "gs/lib/rpc"
	rpclogic "gs/server_node/logic/rpc"
)

type LogicServer struct {
}

func (this *LogicServer) Init() {

}

func (this *LogicServer) Run() {
	fmt.Println("logic server run")
	myrpc.Get().SetName(define.NodeLogic)
	myrpc.Get().NewRpcServer()
	myrpc.Get().RegisterRpcFunc(new(rpclogic.RpcLogic))

}

func (this *LogicServer) Destory() {

}
