package player_info

import (
	"encoding/json"
	"gs/lib/mylog"
	"gs/lib/myredis"
	"gs/lib/myrpc"
	"gs/proto/myproto"
	"sync"
)

type Player struct {
	Uid        uint64
	Name       string
	CreateAt   int64
	NotifyAddr string
	lock       sync.RWMutex
}

func (this *Player) Proto() *myproto.PlayerInfo {
	return &myproto.PlayerInfo{
		Uid:  this.Uid,
		Name: this.Name,
	}
}

func (this *Player) SetNotifyAddr(addr string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.NotifyAddr = addr
}

func (this *Player) Save() bool {
	jsonData, err := json.Marshal(this)
	if err != nil {
		mylog.Error(err)
		return false
	}
	if ok := myredis.GetInstance().Set(myredis.GetRoleKey(this.Uid), jsonData, 0); !ok {
		mylog.Error("player save to redis failed")
		return false
	}
	return true
}

func (this *Player) SendMsg(msgid myproto.MsgId, msg myproto.MyMsg) {
	if this.NotifyAddr == "" {
		return
	}
	data, err := msg.Marshal()
	if err != nil {
		mylog.Error("msg marshal err,msgid ", msgid, ",err:", err)
		return
	}
	myrpc.GetInstance().SendMsg(this.NotifyAddr, this.Uid, msgid, data)
}
