package myrpc

import (
	"encoding/binary"
	"io"

	"github.com/xfxdev/xtcp"
)

type RpcPacket struct {
	Data string
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
	length := len(packet.Data) + 4
	buf := make([]byte, length)
	binary.LittleEndian.PutUint32(buf, uint32(length))
	copy(buf[4:], []byte(packet.Data))
	return buf, nil
}

func (this *RpcProtocol) Unpack(buf []byte) (xtcp.Packet, int, error) {
	length := binary.LittleEndian.Uint32(buf)
	data := buf[4:]
	return &RpcPacket{Data: string(data)}, int(length), nil
}
