package myrpc

import (
	"encoding/binary"
	"io"

	"github.com/xfxdev/xtcp"
)

type RpcPacket struct {
	Uid   uint64
	MsgId uint32
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
	length := len(packet.Data) + 16
	buf := make([]byte, length)
	binary.LittleEndian.PutUint32(buf, uint32(length))
	binary.LittleEndian.PutUint64(buf[4:], packet.Uid)
	binary.LittleEndian.PutUint32(buf[12:], packet.MsgId)
	copy(buf[16:], packet.Data)
	return buf, nil
}

func (this *RpcProtocol) Unpack(buf []byte) (xtcp.Packet, int, error) {
	length := binary.LittleEndian.Uint32(buf)
	uid := binary.LittleEndian.Uint64(buf[4:])
	msgId := binary.LittleEndian.Uint32(buf[12:])
	data := buf[16:]
	return &RpcPacket{Uid: uid, MsgId: msgId, Data: data}, int(length), nil
}
