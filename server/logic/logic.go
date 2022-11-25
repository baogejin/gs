package logic

import (
	"fmt"
	"gs/define"
	"gs/lib/mylog"
	"gs/lib/myrpc"
	rpc_logic "gs/server/logic/rpc"
)

type LogicServer struct {
}

func (this *LogicServer) Init() {

}

func (this *LogicServer) Run() {
	fmt.Println("logic server run")
	port := myrpc.GetInstance().NewRpcServer(define.NodeLogic)
	mylog.Info("rpc server port:", port)
	myrpc.GetInstance().Register(new(rpc_logic.RpcLogic))
}

func (this *LogicServer) Destory() {

}
