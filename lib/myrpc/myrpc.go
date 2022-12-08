package myrpc

import (
	"errors"
	"fmt"
	"gs/lib/myconfig"
	"gs/lib/mylog"
	"gs/lib/myredis"
	"gs/lib/myutil"
	"gs/proto/myproto"
	"net"
	"sync"
	"time"

	"github.com/xfxdev/xtcp"
)

type MyRpc struct {
	name    string
	address string
	server  *Server
	cliMgrs sync.Map
	notifys sync.Map
}

type RpcParam struct {
	Node   string
	Module string
	Fn     string
	Req    RpcReq
	Ack    interface{}
}

type RpcReq interface {
	SetAddr(string)
}
type RpcBaseReq struct {
	Addr string
}

func (this *RpcBaseReq) SetAddr(addr string) {
	this.Addr = addr
}

var myRpc *MyRpc
var once sync.Once

func GetInstance() *MyRpc {
	once.Do(func() {
		myRpc = new(MyRpc)
		myRpc.init()
	})
	return myRpc
}

func (this *MyRpc) init() {
}

func (this *MyRpc) NewRpcServer(name string) string {
	this.name = name
	this.server = NewServer()
	wait := make(chan bool, 1)
	port := myconfig.Get().RpcPortStart
	go func() {
		for {
			address := fmt.Sprintf("%s:%d", myutil.GetLocalIP(), port)
			ln, err := net.Listen("tcp", address)
			if err != nil {
				port++
				time.Sleep(10 * time.Millisecond)
				continue
			}
			this.address = address
			wait <- true
			this.server.Serve(ln)
			return
		}
	}()
	<-wait
	mylog.Info("node ", this.name, " rpc server start address:", this.address)
	return this.address
}

func (this *MyRpc) RegisterFunc(rcvr interface{}) error {
	if this.server == nil {
		return errors.New("server is nil,need NewRpcServer first")
	}
	return this.server.Register(rcvr)
}

func (this *MyRpc) RegisterClient(node string, selector Selector, notifyFn func(p *RpcPacket)) {
	if _, ok := this.cliMgrs.Load(node); ok {
		mylog.Warning("node ", node, " client already register")
		return
	}
	cliMgr := NewClientMgr(node, selector, notifyFn)
	this.cliMgrs.Store(node, cliMgr)
}

func (this *MyRpc) RegisterServerToRedis() {
	if this.server == nil {
		mylog.Warning("rpc server is nil,register to redis failed")
		return
	}
	if this.name == "" {
		mylog.Warning("node name empty,register to redis failed")
		return
	}
	if this.address == "" {
		mylog.Warning("node address empty,register to redis failed")
		return
	}
	myredis.GetInstance().HSet(this.name, this.address, time.Now().Unix())
	myredis.GetInstance().Publish(this.name, nil)
}

func (this *MyRpc) Destory() {
	if this.server != nil {
		myredis.GetInstance().HDel(this.name, this.address)
		myredis.GetInstance().Publish(this.name, nil)
	}
}

func (this *MyRpc) Call(param *RpcParam) (interface{}, error) {
	if cliMgr, ok := this.cliMgrs.Load(param.Node); ok {
		c := cliMgr.(*ClientMgr)
		return c.Call(param)
	}
	return nil, errors.New("node " + param.Node + " not found,need register client")
}

func (this *MyRpc) SendMsg(addr string, uid uint64, msgid myproto.MsgId, data []byte) {
	conn := this.getNotifyConn(addr)
	if conn != nil {
		_, err := conn.SendPacket(&RpcPacket{
			Uid:   uid,
			MsgId: uint32(msgid),
			Data:  data,
		})
		if err != nil {
			mylog.Error("send msg err:", err)
		}
	}
}

func (this *MyRpc) getNotifyConn(addr string) *xtcp.Conn {
	conn, ok := this.notifys.Load(addr)
	if ok {
		client := conn.(*xtcp.Conn)
		if client.IsStoped() {
			this.notifys.Delete(addr)
		} else {
			return client
		}
	}
	option := xtcp.NewOpts(&RpcHandler{}, &RpcProtocol{}) //这个连接只发送，不处理收到的消息，故不设置处理函数
	// option.SendBufListLen = 4096
	client := xtcp.NewConn(option)
	this.notifys.Store(addr, client)
	go func() {
		err := client.DialAndServe(addr)
		if err != nil {
			mylog.Error(err)
			this.notifys.Delete(addr)
		}
	}()
	return client
}
