package servernode

import (
	"gs/define"
	gateway "gs/serverNode/gateway"
	logic "gs/serverNode/logic"
)

type MyServerNode interface {
	Init()
	Run()
	Destory()
}

func GetServerNode(name string) MyServerNode {
	switch name {
	case define.NodeGateway:
		return new(gateway.GatewayServer)
	case define.NodeLogic:
		return new(logic.LogicServer)
	}
	return nil
}
