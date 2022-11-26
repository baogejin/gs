package logic

import (
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
	mylog.Info("logic server run")
	myrpc.GetInstance().NewRpcServer(define.NodeLogic)
	myrpc.GetInstance().RegisterFunc(new(rpc_logic.RpcLogic))
	myrpc.GetInstance().RegisterServerToRedis()
}

func (this *LogicServer) Destory() {
	myrpc.GetInstance().Destory()
}
