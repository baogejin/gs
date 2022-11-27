package myrpc

import (
	"github.com/xfxdev/xtcp"
)

type RpcHandler struct {
	notifyFn func(p *RpcPacket)
}

func NewRpcHandler(fn func(p *RpcPacket)) *RpcHandler {
	return &RpcHandler{
		notifyFn: fn,
	}
}

func (this *RpcHandler) OnAccept(conn *xtcp.Conn) {

}

func (this *RpcHandler) OnConnect(conn *xtcp.Conn) {

}

func (this *RpcHandler) OnRecv(conn *xtcp.Conn, p xtcp.Packet) {
	if this.notifyFn != nil {
		packet := p.(*RpcPacket)
		this.notifyFn(packet)
	}
}

func (this *RpcHandler) OnUnpackErr(c *xtcp.Conn, buf []byte, err error) {

}

func (this *RpcHandler) OnClose(conn *xtcp.Conn) {

}
