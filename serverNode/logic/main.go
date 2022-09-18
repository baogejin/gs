package logic

import (
	"fmt"
	"gs/define"
	myrpc "gs/lib/myRpc"
	rpclogic "gs/serverNode/logic/rpc"
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
