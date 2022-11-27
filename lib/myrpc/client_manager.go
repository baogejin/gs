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

	"github.com/xfxdev/xtcp"
)

type ClientMgr struct {
	node       string
	selector   Selector
	clients    sync.Map
	needUpdate bool
	lock       sync.RWMutex
	xtcpServer *xtcp.Server
	address    string
}

func NewClientMgr(node string, selector Selector) *ClientMgr {
	if selector == nil {
		selector = &RoundSelector{}
	}
	option := xtcp.NewOpts(&RpcHandler{}, &RpcProtocol{})
	option.SendBufListLen = 4096
	s := xtcp.NewServer(option)
	mgr := &ClientMgr{
		node:       node,
		selector:   selector,
		needUpdate: true,
		xtcpServer: s,
	}
	mgr.init()
	return mgr
}

func (this *ClientMgr) init() {
	port := myconfig.Get().RpcPortStart
	wait := make(chan bool, 1)
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
			this.xtcpServer.Serve(ln)
			return
		}
	}()
	<-wait
	mylog.Info("rpc xtcp address ", this.address)

	ch := myredis.GetInstance().Subscribe(this.node).Channel()
	go func() {
		for {
			select {
			case <-ch:
				this.lock.Lock()
				this.needUpdate = true
				this.lock.Unlock()
			}
		}
	}()
}

func (this *ClientMgr) Call(param *RpcParam) (interface{}, error) {
	if param == nil {
		return nil, errors.New("param is nil")
	}
	if this.needUpdate {
		this.lock.Lock()
		this.needUpdate = false
		this.lock.Unlock()
		servers := myredis.GetInstance().HGetAll(this.node)
		needDelete := []string{}
		this.clients.Range(func(key, value any) bool {
			addr := fmt.Sprintf("%v", key)
			if _, ok := servers[addr]; !ok {
				needDelete = append(needDelete, addr)
				cli := value.(*Client)
				cli.Close()
			}
			return true
		})
		for _, v := range needDelete {
			this.clients.Delete(v)
		}
		this.selector.UpdateServer(servers)
	}
	addr := this.selector.Select(param.Req)
	if addr == "" {
		return nil, errors.New("select " + this.node + " rpc addr is empty")
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
	param.Req.SetAddr(this.address)

	err := c.Call(fmt.Sprintf("%s.%s", param.Module, param.Fn), param.Req, param.Ack)
	if err != nil {
		return nil, err
	}
	return param.Ack, nil
}
