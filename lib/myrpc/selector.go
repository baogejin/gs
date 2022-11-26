package myrpc

import (
	"gs/lib/myredis"
	"math/rand"
	"sync"
	"time"
)

type Selector interface {
	SetNode(node string)
	Select(req interface{}) string
}

type ServerInfo struct {
	Address string
	Info    string
}

type DefaultSelector struct {
	node     string
	servers  []*ServerInfo
	updateAt int64
	lock     sync.RWMutex
}

func (this *DefaultSelector) SetNode(node string) {
	this.node = node
}

func (this *DefaultSelector) Select(req interface{}) string {
	if len(this.servers) == 0 || time.Now().Unix()-this.updateAt > 10 {
		this.updateServers()
	}
	if len(this.servers) == 0 {
		return ""
	}
	idx := rand.Int31n(int32(len(this.servers)))
	return this.servers[idx].Address
}

func (this *DefaultSelector) updateServers() {
	this.servers = this.servers[:0]
	servers := myredis.GetInstance().HGetAll(this.node)
	for k, v := range servers {
		this.servers = append(this.servers, &ServerInfo{Address: k, Info: v})
	}
	this.updateAt = time.Now().Unix()
}
