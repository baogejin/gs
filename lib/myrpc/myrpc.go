package myrpc

import (
	"errors"
	"fmt"
	"gs/lib/myconfig"
	"gs/lib/mylog"
	"gs/lib/myredis"
	"gs/lib/myutil"
	"net"
	"sync"
	"time"
)

type MyRpc struct {
	server    *Server
	clients   sync.Map
	name      string
	address   string
	selectors map[string]Selector
}

type RpcParam struct {
	Node   string
	Module string
	Fn     string
	Req    interface{}
	Ack    interface{}
}

var myRpc *MyRpc
var once sync.Once

func GetInstance() *MyRpc {
	once.Do(func() {
		myRpc = new(MyRpc)
		myRpc.init()
	})
	return myRpc
}

func (this *MyRpc) init() {
	this.selectors = make(map[string]Selector)
}

func (this *MyRpc) NewRpcServer(name string) string {
	this.name = name
	this.server = NewServer()
	wait := make(chan bool, 1)
	port := myconfig.Get().RpcPortStart
	go func() {
		for {
			address := fmt.Sprintf("%s:%d", myutil.GetLocalIP(), port)
			ln, err := net.Listen("tcp", address)
			if err != nil {
				port++
				time.Sleep(10 * time.Millisecond)
				continue
			}
			this.address = address
			wait <- true
			this.server.Serve(ln)
			return
		}
	}()
	<-wait
	mylog.Info("node ", this.name, " rpc server start address:", this.address)
	return this.address
}

func (this *MyRpc) RegisterFunc(rcvr interface{}) error {
	if this.server == nil {
		return errors.New("server is nil,need NewRpcServer first")
	}
	return this.server.Register(rcvr)
}

func (this *MyRpc) RegisterClient(node string, selector Selector) {
	if _, ok := this.selectors[node]; ok {
		mylog.Warning("node ", node, " client already register")
		return
	}
	if selector == nil {
		selector = &DefaultSelector{}
	}
	selector.SetNode(node)
	this.selectors[node] = selector
}

func (this *MyRpc) RegisterServerToRedis() {
	if this.server == nil {
		mylog.Warning("rpc server is nil,register to redis failed")
		return
	}
	if this.name == "" {
		mylog.Warning("node name empty,register to redis failed")
		return
	}
	if this.address == "" {
		mylog.Warning("node address empty,register to redis failed")
		return
	}
	myredis.GetInstance().HSet(this.name, this.address, time.Now().Unix())
}

func (this *MyRpc) Destory() {
	if this.server != nil {
		myredis.GetInstance().HDel(this.name, this.address)
	}
}

func (this *MyRpc) Call(param *RpcParam) (interface{}, error) {
	if param == nil {
		return nil, errors.New("param is nil")
	}
	selector := this.selectors[param.Node]
	if selector == nil {
		return nil, errors.New("can not find selector " + param.Node)
	}
	addr := selector.Select(param.Req)
	if addr == "" {
		return nil, errors.New("select rpc addr is empty")
	}
	var c *Client
	if cli, ok := this.clients.Load(addr); ok {
		c = cli.(*Client)
	} else {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			mylog.Error(err)
			return nil, err
		}
		c = NewClient(conn)
		this.clients.Store(addr, c)
	}
	err := c.Call(fmt.Sprintf("%s.%s", param.Module, param.Fn), param.Req, param.Ack)
	if err != nil {
		return nil, err
	}
	return param.Ack, nil
}
