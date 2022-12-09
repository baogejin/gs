package gateway

import (
	"gs/lib/myrpc"
	rpc_logic "gs/server/logic/rpc"
	"math/rand"
	"sync"
)

type LogicSelector struct {
	servers []*myrpc.ServerInfo
	addrMap map[string]bool
	records map[uint64]string
	lock    sync.RWMutex
}

func (this *LogicSelector) Select(req interface{}) string {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if len(this.servers) == 0 {
		return ""
	}
	logicReq, ok := req.(*rpc_logic.LogicReq)
	if !ok {
		return ""
	}
	if logicReq.Uid == 0 {
		idx := rand.Int31n(int32(len(this.servers)))
		return this.servers[idx].Address
	}
	addr, ok := this.records[logicReq.Uid]
	if ok {
		if this.addrMap != nil && this.addrMap[addr] {
			return addr
		} else {
			return ""
		}
	}
	idx := rand.Int31n(int32(len(this.servers)))
	addr = this.servers[idx].Address
	this.records[logicReq.Uid] = addr
	return addr
}

func (this *LogicSelector) UpdateServer(servers map[string]string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.servers = this.servers[:0]
	this.addrMap = make(map[string]bool)
	for k, v := range servers {
		this.servers = append(this.servers, &myrpc.ServerInfo{Address: k, Info: v})
		this.addrMap[k] = true
	}
}
