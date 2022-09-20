package server_node

import (
	"gs/define"
	gateway "gs/server_node/gateway"
	logic "gs/server_node/logic"
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
