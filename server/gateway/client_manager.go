package gateway

import "sync"

type ClientMgr struct {
	clients sync.Map
}

var cliMgr *ClientMgr
var cliMgrOnce sync.Once

func GetClinetMgr() *ClientMgr {
	cliMgrOnce.Do(func() {
		cliMgr = new(ClientMgr)
		cliMgr.init()
	})
	return cliMgr
}

func (this *ClientMgr) init() {

}

func (this *ClientMgr) AddClient(uid uint64, cli *Client) {
	c, ok := this.clients.Load(uid)
	if ok {
		client := c.(*Client)
		client.Kick()
		this.clients.Delete(uid)
	}
	this.clients.Store(uid, cli)
}
