package myrpc

import (
	"gs/lib/mylog"

	"github.com/xfxdev/xtcp"
)

type RpcHandler struct {
}

func (this *RpcHandler) OnAccept(conn *xtcp.Conn) {

}

func (this *RpcHandler) OnConnect(conn *xtcp.Conn) {

}

func (this *RpcHandler) OnRecv(conn *xtcp.Conn, p xtcp.Packet) {
	packet := p.(*RpcPacket)
	mylog.Info(string(packet.Data))
}

func (this *RpcHandler) OnUnpackErr(c *xtcp.Conn, buf []byte, err error) {

}

func (this *RpcHandler) OnClose(conn *xtcp.Conn) {

}
