package gateway

import (
	"encoding/binary"
	"gs/define"
	"gs/lib/mylog"
	"gs/lib/myrpc"
	"gs/proto/myproto"
	rpc_logic "gs/server/logic/rpc"
	"io"

	"golang.org/x/net/websocket"
)

type ClientConn struct {
	ws     *websocket.Conn
	buf    []byte
	bufLen uint32
	seq    uint64
	uid    uint64
}

func (this *ClientConn) Start() {
	defer func() {
		mylog.Info("ws conn close")
		this.ws.Close()
	}()
	this.buf = make([]byte, 2048)
	this.bufLen = 0
	this.seq = 0
	for {
		length, err := this.ws.Read(this.buf[this.bufLen:])
		if err == io.EOF {
			return
		}
		if err != nil {
			mylog.Info("ws read err:", err)
			return
		}
		if length == 0 {
			mylog.Info("消息超过了缓冲长度")
			return
		}
		this.bufLen += uint32(length)
		if this.bufLen > 4 {
			needLen := binary.LittleEndian.Uint32(this.buf)
			if this.bufLen >= needLen {
				msg := UnpackMsg(this.buf[4:needLen])
				this.seq++
				if this.seq != msg.Seq {
					//TODO 序列不对
				}
				this.ProcessMsg(msg.MsgId, msg.Data)
				this.buf = this.buf[needLen:]
				this.bufLen -= needLen
			}
		}
	}
}

func (this *ClientConn) ProcessMsg(msgId uint32, data []byte) {
	if this.uid == 0 && msgId != uint32(myproto.MsgId_Msg_RegisterREQ) && msgId != uint32(myproto.MsgId_Msg_LoginREQ) {
		//需要先登录
	} else {
		reply, err := myrpc.GetInstance().Call(&myrpc.RpcParam{
			Node:   define.NodeLogic,
			Module: "RpcLogic",
			Fn:     "Logic",
			Req: &rpc_logic.LogicReq{
				Uid:   this.uid,
				MsgId: msgId,
				Data:  data,
			},
			Ack: &rpc_logic.LogicAck{},
		})
		if err != nil {
			//rpc 错误
		} else {
			ack := reply.(*rpc_logic.LogicAck)
			if ack.MsgId == uint32(myproto.MsgId_Msg_LoginACK) {
				msg := &myproto.LoginACK{}
				if err := msg.Unmarshal(ack.Data); err == nil {
					if msg.Ret == myproto.ResultCode_Success && this.uid == 0 {
						this.uid = msg.Uid
					}
				}
			}
			this.ws.Write(PackMsg(ack.MsgId, ack.Data))
		}
	}
}
