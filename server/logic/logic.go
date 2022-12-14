package logic

import (
	"gs/define"
	"gs/lib/mylog"
	"gs/lib/myrpc"
	"gs/proto/myproto"
	"gs/server/logic/player_manager"
	rpc_logic "gs/server/logic/rpc"
)

type LogicServer struct {
}

func (this *LogicServer) Init() {

}

func (this *LogicServer) Run() {
	mylog.Info("logic server run")
	myrpc.GetInstance().SetNodeName(define.NodeLogic)
	myrpc.GetInstance().SetNotifyHandler(this.handleNotify)
	myrpc.GetInstance().NewRpcServer()
	myrpc.GetInstance().RegisterFunc(new(rpc_logic.RpcLogic))
	myrpc.GetInstance().RegisterRpcServerToRedis()
}

func (this *LogicServer) Destory() {
	player_manager.GetMgr().Destory()
	myrpc.GetInstance().Destory()
}

func (this *LogicServer) handleNotify(p *myrpc.RpcPacket) {
	if p.Node != define.NodeId[define.NodeLogic] {
		return
	}
	player := player_manager.GetMgr().GetPlayer(p.Uid)
	if player != nil {
		player.ProcessNotify(myproto.MsgId(p.MsgId), p.Data)
	}
}
