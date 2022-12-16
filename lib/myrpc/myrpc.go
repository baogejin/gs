package myrpc

import (
	"errors"
	"fmt"
	"gs/define"
	"gs/lib/myconfig"
	"gs/lib/mylog"
	"gs/lib/myredis"
	"gs/lib/myticker"
	"gs/lib/myutil"
	"gs/proto/myproto"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/xfxdev/xtcp"
)

const (
	UpdateTime = time.Second * 10
	CheckTime  = time.Second * 12
)

type MyRpc struct {
	name       string
	address    string
	server     *Server
	cliMgrs    sync.Map
	notifys    sync.Map
	notifySvr  *xtcp.Server
	notifyAddr string
}

type RpcParam struct {
	Node   string
	Module string
	Fn     string
	Req    RpcReq
	Ack    interface{}
}

type RpcReq interface {
	SetNotifyAddr(string)
}
type RpcBaseReq struct {
	NotifyAddr string
}

func (this *RpcBaseReq) SetNotifyAddr(addr string) {
	this.NotifyAddr = addr
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

func (this *MyRpc) SetNodeName(name string) {
	this.name = name
}

func (this *MyRpc) NewRpcServer() string {
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

func (this *MyRpc) SetNotifyHandler(fn func(p *RpcPacket)) {
	handler := NewRpcHandler(fn)
	option := xtcp.NewOpts(handler, &RpcProtocol{})
	option.SendBufListLen = 4096
	this.notifySvr = xtcp.NewServer(option)
	port := myconfig.Get().RpcPortStart
	wait := make(chan bool, 1)
	go func() {
		for {
			address := fmt.Sprintf("%s:%d", myutil.GetLocalIP(), port)
			ln, err := net.Listen("tcp", address)
			if err != nil {
				port++
				time.Sleep(10 * time.Millisecond)
				continue
			}
			this.notifyAddr = address
			wait <- true
			this.notifySvr.Serve(ln)
			return
		}
	}()
	<-wait
	mylog.Info("rpc notify address ", this.notifyAddr)
	myredis.GetInstance().HSet("notify_"+this.name, this.notifyAddr, time.Now().Unix())
	//定时更新
	myticker.GetInstance().AddTicker(UpdateTime, func() {
		myredis.GetInstance().HSet("notify_"+this.name, this.notifyAddr, time.Now().Unix())
	})
}

func (this *MyRpc) GetNotifyAddr() string {
	return this.notifyAddr
}

func (this *MyRpc) RegisterFunc(rcvr interface{}) error {
	if this.server == nil {
		mylog.Warning("server is nil,need NewRpcServer first")
		return errors.New("server is nil,need NewRpcServer first")
	}
	return this.server.Register(rcvr)
}

func (this *MyRpc) RegisterClient(node string, selector Selector) {
	if _, ok := this.cliMgrs.Load(node); ok {
		mylog.Warning("node ", node, " client already register")
		return
	}
	cliMgr := NewClientMgr(node, selector)
	this.cliMgrs.Store(node, cliMgr)
}

func (this *MyRpc) RegisterRpcServerToRedis() {
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
	myticker.GetInstance().AddTicker(UpdateTime, func() {
		myredis.GetInstance().HSet(this.name, this.address, time.Now().Unix())
	})
}

func (this *MyRpc) Destory() {
	if this.server != nil {
		myredis.GetInstance().HDel(this.name, this.address)
		myredis.GetInstance().Publish(this.name, nil)
	}
	if this.notifySvr != nil {
		myredis.GetInstance().HDel("notify_"+this.name, this.notifyAddr)
	}
}

func (this *MyRpc) Call(param *RpcParam) (interface{}, error) {
	if cliMgr, ok := this.cliMgrs.Load(param.Node); ok {
		c := cliMgr.(*ClientMgr)
		param.Req.SetNotifyAddr(this.notifyAddr)
		return c.Call(param)
	}
	return nil, errors.New("node " + param.Node + " not found,need register client")
}

func (this *MyRpc) SendMsg(addr string, uid uint64, msgid myproto.MsgId, node string, data []byte) {
	conn := this.getNotifyConn(addr)
	if conn != nil {
		_, err := conn.SendPacket(&RpcPacket{
			Uid:   uid,
			MsgId: uint32(msgid),
			Node:  define.NodeId[node],
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

func (this *MyRpc) NotifyAllNodes(node string, msgid myproto.MsgId, msg myproto.MyMsg) {
	data, err := msg.Marshal()
	if err != nil {
		mylog.Error(err)
		return
	}
	nodes := myredis.GetInstance().HGetAll("notify_" + node)
	for k, v := range nodes {
		updateAt, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			mylog.Error(err)
			continue
		}
		if time.Now().Unix()-updateAt > int64(CheckTime/time.Second) {
			continue
		}
		this.SendMsg(k, 0, msgid, node, data)
	}
}
