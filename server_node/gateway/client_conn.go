package gateway

import (
	"encoding/binary"
	"fmt"
	"gs/define"
	"gs/lib/myrpc"
	rpc_logic "gs/server_node/logic/rpc"
	"io"

	"golang.org/x/net/websocket"
)

type ClientConn struct {
	ws     *websocket.Conn
	buf    []byte
	bufLen uint32
	seq    uint64
}

func (this *ClientConn) Start() {
	defer func() {
		fmt.Println("ws conn close")
		this.ws.Close()
	}()
	this.buf = make([]byte, 2048)
	this.bufLen = 0
	this.seq = 0
	for {
		length, err := this.ws.Read(this.buf[this.bufLen:])
		if err == io.EOF {
			fmt.Println("ws conn EOF")
			return
		}
		if err != nil {
			fmt.Println("ws read err:", err)
			return
		}
		if length == 0 {
			fmt.Println("消息超过了缓冲长度")
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
					fmt.Println("seq err")
				}
				this.ProcessMsg(msg.MsgId, msg.Data)
				this.buf = this.buf[needLen:]
				this.bufLen -= needLen
			}
		}
	}
}

func (this *ClientConn) ProcessMsg(msgId uint32, data []byte) {
	fmt.Printf("%d,%s\n", msgId, data)
	ret, err := myrpc.Get().RpcRun(&myrpc.RpcParm{
		Node:      define.NodeLogic,
		RpcModule: "RpcLogic",
		Fn:        "Logic",
		Arg:       &rpc_logic.LogicReq{MsgId: msgId, Data: data},
		Reply:     &rpc_logic.LogicAck{},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	ack := ret.(*rpc_logic.LogicAck)
	fmt.Printf("ack:%d,%s\n", ack.MsgId, ack.Data)
}
