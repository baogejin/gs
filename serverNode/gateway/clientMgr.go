package gateway

import (
	"sync"

	"golang.org/x/net/websocket"
)

type ClientMgr struct {
	clientId  uint64
	clientMap sync.Map
}

func (this *ClientMgr) OnNewConn(conn *websocket.Conn) {

}

func (this *ClientMgr) OnCloseConn(id uint64) {

}
