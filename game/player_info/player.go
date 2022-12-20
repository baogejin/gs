package player_info

import (
	"encoding/json"
	"gs/define"
	"gs/game/battle"
	"gs/game/player_info/bag"
	"gs/lib/mylog"
	"gs/lib/myredis"
	"gs/lib/myrpc"
	"gs/proto/myproto"
	"sync"
	"time"
)

type Player struct {
	Uid        uint64       `json:"Uid,,omitempty"`
	Name       string       `json:"Name,omitempty"`
	CreateAt   int64        `json:"CreateAt,,omitempty"`
	NotifyAddr string       `json:"-"`
	lock       sync.RWMutex `json:"-"`
	Bag        *bag.Bag     `json:"Bag,,omitempty"`
}

func NewPlayer(uid uint64, name string) *Player {
	return &Player{
		Uid:      uid,
		Name:     name,
		CreateAt: time.Now().Unix(),
		Bag:      bag.NewBag(),
	}
}

func (this *Player) Proto() *myproto.PlayerInfo {
	return &myproto.PlayerInfo{
		Uid:     this.Uid,
		Name:    this.Name,
		BagInfo: this.Bag.Proto(),
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
	myrpc.GetInstance().SendMsg(this.NotifyAddr, this.Uid, msgid, define.NodeGateway, data)
}

func (this *Player) SendMsgData(msgid myproto.MsgId, data []byte) {
	if this.NotifyAddr == "" {
		return
	}
	myrpc.GetInstance().SendMsg(this.NotifyAddr, this.Uid, msgid, define.NodeGateway, data)
}

func (this *Player) ProcessNotify(msgId myproto.MsgId, data []byte) {
	switch msgId {
	default:
		this.SendMsgData(msgId, data)
	}
}

//临时生成战斗单位
func (this *Player) GenBattleUnit() *battle.Unit {
	return &battle.Unit{
		Uid:          this.Uid,
		Name:         this.Name,
		Team:         0,
		Position:     2,
		UnitType:     myproto.UnitType_UnitPlayer,
		WeaponSkill:  1,
		SkillUseTime: make(map[int32]int64),
		HP:           100,
		MaxHP:        100,
		NotifyAddr:   this.NotifyAddr,
	}
}
