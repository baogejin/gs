package myrpc

import (
	"errors"
	"fmt"
	"gs/lib/mylog"
	"gs/lib/myredis"
	"net"
	"strconv"
	"sync"
	"time"
)

type ClientMgr struct {
	node       string
	selector   Selector
	clients    sync.Map
	needUpdate bool
	lastUpdate int64
	lock       sync.RWMutex
}

func NewClientMgr(node string, selector Selector) *ClientMgr {
	if selector == nil {
		selector = &RoundSelector{}
	}
	mgr := &ClientMgr{
		node:       node,
		selector:   selector,
		needUpdate: true,
	}
	mgr.init()
	return mgr
}

func (this *ClientMgr) init() {
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
	if this.needUpdate || time.Now().Unix()-this.lastUpdate >= int64(CheckTime/time.Second) {
		this.lock.Lock()
		this.needUpdate = false
		this.lastUpdate = time.Now().Unix()
		this.lock.Unlock()
		s := myredis.GetInstance().HGetAll(this.node)
		servers := make(map[string]string)
		for k, v := range s {
			//过滤掉长久不更新的记录
			updateAt, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				mylog.Error(err)
				continue
			}
			if time.Now().Unix()-updateAt > int64(CheckTime/time.Second) {
				continue
			}
			servers[k] = v
		}
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

	err := c.Call(fmt.Sprintf("%s.%s", param.Module, param.Fn), param.Req, param.Ack)
	if err != nil {
		return nil, err
	}
	return param.Ack, nil
}
