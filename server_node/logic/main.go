package logic

import (
	"fmt"
	"gs/define"
	"gs/lib/myrpc"
	rpc_logic "gs/server_node/logic/rpc"
)

type LogicServer struct {
}

func (this *LogicServer) Init() {

}

func (this *LogicServer) Run() {
	fmt.Println("logic server run")
	myrpc.Get().Init(define.NodeLogic, []string{"127.0.0.1:8500"})
	port, err := myrpc.Get().NewRpcServer()
	if err != nil {
		fmt.Println("NewRpcServer failed,", err)
		return
	}
	fmt.Println("rpc server port:", port)
	myrpc.Get().RegisterRpcFunc(new(rpc_logic.RpcLogic))

}

func (this *LogicServer) Destory() {

}
