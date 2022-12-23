package myrpc

import (
	"encoding/binary"
	"io"

	"github.com/xfxdev/xtcp"
)

type RpcPacket struct {
	Uid   uint64
	MsgId uint32
	Node  int32
	Data  []byte
}

func (this *RpcPacket) String() string {
	return ""
}

type RpcProtocol struct {
}

func (this *RpcProtocol) PackSize(p xtcp.Packet) int {
	return 4
}

func (this *RpcProtocol) PackTo(p xtcp.Packet, w io.Writer) (int, error) {
	return 0, nil
}

func (this *RpcProtocol) Pack(p xtcp.Packet) ([]byte, error) {
	packet := p.(*RpcPacket)
	length := len(packet.Data) + 20
	buf := make([]byte, length)
	binary.LittleEndian.PutUint32(buf, uint32(length))
	binary.LittleEndian.PutUint64(buf[4:], packet.Uid)
	binary.LittleEndian.PutUint32(buf[12:], packet.MsgId)
	binary.LittleEndian.PutUint32(buf[16:], uint32(packet.Node))
	copy(buf[20:], packet.Data)
	return buf, nil
}

func (this *RpcProtocol) Unpack(buf []byte) (xtcp.Packet, int, error) {
	length := binary.LittleEndian.Uint32(buf)
	uid := binary.LittleEndian.Uint64(buf[4:])
	msgId := binary.LittleEndian.Uint32(buf[12:])
	node := binary.LittleEndian.Uint32(buf[16:])
	data := buf[20:length]
	return &RpcPacket{Uid: uid, MsgId: msgId, Node: int32(node), Data: data}, int(length), nil
}
