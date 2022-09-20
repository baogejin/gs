package gateway

import "encoding/binary"

type MsgPack struct {
	Seq   uint64
	MsgId uint32
	Data  []byte
}

func PackMsg(msgId uint32, data []byte) []byte {
	length := 4 + 8 + 4 + len(data)
	buf := make([]byte, length)
	binary.LittleEndian.PutUint32(buf, uint32(length))
	binary.LittleEndian.PutUint64(buf[4:], 0)
	binary.LittleEndian.PutUint32(buf[12:], msgId)
	copy(buf[16:], data)
	return buf
}

func UnpackMsg(buf []byte) *MsgPack {
	msg := &MsgPack{}
	msg.Seq = binary.LittleEndian.Uint64(buf)
	msg.MsgId = binary.LittleEndian.Uint32(buf[8:])
	msg.Data = buf[12:]
	return msg
}
