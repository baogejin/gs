package main

import (
	"flag"
	"fmt"
	"gs/server_node"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		nodeName string
	)
	flag.StringVar(&nodeName, "node", "gateway", "server node")
	flag.Parse()

	server := server_node.GetServerNode(nodeName)
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
