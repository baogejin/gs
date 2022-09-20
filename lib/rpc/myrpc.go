package myrpc

import (
	"context"
	"errors"
	"fmt"
	myutil "gs/lib/util"
	"net"
	"sync"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

type MyRpc struct {
	Name       string
	server     *server.Server
	clients    sync.Map
	Address    string
	Port       int32
	Exits      []chan bool
	clientExit sync.WaitGroup
}

var once sync.Once
var myRpcInstance *MyRpc

func Get() *MyRpc {
	once.Do(func() {
		myRpcInstance = &MyRpc{}
	})
	return myRpcInstance
}

func (this *MyRpc) SetName(name string) {
	this.Name = name
}

func (this *MyRpc) NewRpcServer() int32 {
	if this.server != nil {
		return 0
	}
	s := server.NewServer()
	wait := make(chan bool, 1)
	localIp := myutil.GetLocalIP()
	port := int32(6502)
	go func() {
		for {
			address := fmt.Sprintf("%s:%d", localIp, port)
			ln, err := net.Listen("tcp", address)
			if err == nil {
				wait <- true
				err = s.ServeListener("tcp", ln)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
			port++
			time.Sleep(20 * time.Millisecond)
		}

	}()
	<-wait

	//consul
	plug := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: fmt.Sprintf("%s:%d", localIp, port),
		ConsulServers:  []string{"127.0.0.1:8500"},
		BasePath:       fmt.Sprintf("/%s", this.Name),
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: 10 * time.Second,
	}
	plug.Start()
	s.Plugins.Add(plug)
	this.server = s
	this.Address = fmt.Sprintf("tcp@%s:%d", localIp, port)
	this.Port = port

	return port
}

func (this *MyRpc) RegisterRpcFunc(fn interface{}) {
	if this.server == nil {
		return
	}

	this.server.Register(fn, "")
}

func (this *MyRpc) NewRpcClient(nodeName string, rpcName string, selector client.Selector, notifyFn func(msg *protocol.Message)) {
	d, _ := client.NewConsulDiscoveryTemplate(fmt.Sprintf("/%s", nodeName), []string{"127.0.0.1:8500"}, nil)

	svrMsg := make(chan *protocol.Message, 1024)
	option := client.DefaultOption
	option.Retries = 0

	if selector == nil {
		selector = &DefaultSelector{}
	}

	pool := client.NewBidirectionalOneClientPool(2, client.Failover, client.SelectByUser, d, option, svrMsg)
	for i := 0; i < 2; i++ {
		c := pool.Get()
		c.SetSelector(rpcName, selector)
	}
	this.clients.Store(nodeName, pool)
	if notifyFn != nil {
		this.clientExit.Add(1)
		go func() {
			exit := make(chan bool, 1)
			this.Exits = append(this.Exits, exit)
			defer func() {
				log.Info("rpc client ", nodeName, " exit")
				this.clientExit.Done()
			}()
			for {
				select {
				case <-exit:
					return
				case msg := <-svrMsg:
					if notifyFn != nil {
						notifyFn(msg)
					}
				}
			}
		}()
	}
}

func (this *MyRpc) Destory() {
	if this.server != nil {
		this.server.UnregisterAll()
		this.server.Close()
	}
	this.clients.Range(func(k, v interface{}) bool {
		switch v.(type) {
		case *client.OneClientPool:
			pool := v.(*client.OneClientPool)
			pool.Close()
		case map[string]*client.Client:
			clients := v.(map[string]*client.Client)
			for _, v := range clients {
				v.Close()
			}
		}
		return true
	})
	for _, exit := range this.Exits {
		exit <- true
	}

	this.clientExit.Wait()
}

func (this *MyRpc) getClient(nodeName string) *client.OneClientPool {
	if c, ok := this.clients.Load(nodeName); ok {
		return c.(*client.OneClientPool)
	}
	return nil
}

type RpcParm struct {
	Node      string
	RpcModule string
	Fn        string
	Arg       interface{}
	Reply     interface{}
}

func (this *MyRpc) RpcRun(parm *RpcParm) (ret interface{}, err error) {
	pool := this.getClient(parm.Node)
	if pool != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		err := pool.Get().Call(ctx, parm.RpcModule, parm.Fn, parm.Arg, parm.Reply)
		if err != nil {
			fmt.Printf("rpc call err [%s] : %s", parm.Fn, err)
		}
		return parm.Reply, err
	}

	return ret, errors.New("no client")
}
