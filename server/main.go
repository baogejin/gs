package main

import (
	"flag"
	"fmt"
	"gs/define"
	"gs/server/gateway"
	"gs/server/logic"
	"os"
	"os/signal"
	"syscall"
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

func main() {
	var (
		nodeName string
	)
	flag.StringVar(&nodeName, "node", "gateway", "server node")
	flag.Parse()

	server := GetServerNode(nodeName)
	if server == nil {
		fmt.Println("get server node is nil")
		return
	}
	server.Init()
	server.Run()
	exit := make(chan bool, 1)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		sig := <-ch
		fmt.Println("ServerNode:", nodeName, " exit signal:", sig.String())

		exit <- true
	}()
	<-exit

	server.Destory()
}
