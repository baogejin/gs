package gateway

import (
	"bytes"
	"encoding/binary"
	"gs/define"
	"gs/lib/mylog"
	"gs/lib/myrpc"
	"gs/proto/myproto"
	rpc_logic "gs/server/logic/rpc"
	"io"

	"golang.org/x/net/websocket"
)

type Client struct {
	ws  *websocket.Conn
	seq uint64
	uid uint64
}

func (this *Client) Start() {
	defer func() {
		//通知logic下线
		if this.uid > 0 {
			this.ProcessMsg(uint32(myproto.MsgId_Msg_LogoutREQ), []byte{})
			GetClinetMgr().DelClient(this.uid)
		}
		mylog.Info("ws conn close")
		this.ws.Close()
	}()
	buf := make([]byte, 2048)
	recvBuf := bytes.NewBuffer(make([]byte, 0, 2048))
	this.seq = 0
	for {
		length, err := this.ws.Read(buf)
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
		recvBuf.Write(buf[:length])
		for recvBuf.Len() > 4 {
			needLen := binary.LittleEndian.Uint32(recvBuf.Bytes())
			if recvBuf.Len() >= int(needLen) {
				msg := UnpackMsg(recvBuf.Bytes()[4:])
				this.seq++
				if this.seq != msg.Seq {
					//TODO 序列不对
				}
				if !this.ProcessMsg(msg.MsgId, msg.Data) {
					return
				}
				recvBuf.Next(int(needLen))
			} else {
				break
			}
		}
	}
}

func (this *Client) ProcessMsg(msgId uint32, data []byte) bool {
	if this.uid == 0 && msgId != uint32(myproto.MsgId_Msg_RegisterREQ) && msgId != uint32(myproto.MsgId_Msg_LoginREQ) {
		//需要先登录
		return false
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
			mylog.Error(this.uid, " ", err)
			return false
		} else {
			ack := reply.(*rpc_logic.LogicAck)
			if ack.MsgId == uint32(myproto.MsgId_Msg_LoginACK) {
				msg := &myproto.LoginACK{}
				if err := msg.Unmarshal(ack.Data); err == nil {
					if msg.Ret == myproto.ResultCode_Success {
						if this.uid == 0 && msg.Uid > 0 {
							this.uid = msg.Uid
							GetClinetMgr().AddClient(this.uid, this)
						}
					}
				}
			}
			websocket.Message.Send(this.ws, PackMsg(ack.MsgId, ack.Data))
		}
	}
	return true
}

func (this *Client) Kick() {
	websocket.Message.Send(this.ws, PackMsg(uint32(myproto.MsgId_Msg_KickPUSH), []byte{}))
	this.uid = 0
	this.ws.Close()
}

func (this *Client) SendMsg(msgid uint32, data []byte) {
	websocket.Message.Send(this.ws, PackMsg(msgid, data))
}
