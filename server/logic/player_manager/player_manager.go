package player_manager

import (
	"fmt"
	"gs/game/player_info"
	"gs/lib/mylog"
	"gs/lib/myredis"
	"gs/lib/myrpc"
	"gs/proto/myproto"
	"sync"
)

type PlayerMgr struct {
	players sync.Map
}

var playerMgr *PlayerMgr
var once sync.Once

func GetMgr() *PlayerMgr {
	once.Do(func() {
		playerMgr = new(PlayerMgr)
		playerMgr.init()
	})
	return playerMgr
}

func (this *PlayerMgr) init() {

}

func (this *PlayerMgr) GetPlayer(uid uint64) *player_info.Player {
	if p, ok := this.players.Load(uid); ok {
		return p.(*player_info.Player)
	}
	return nil
}

func (this *PlayerMgr) SetPlayer(uid uint64, player *player_info.Player) {
	this.players.Store(uid, player)
	myredis.GetInstance().HSet(myredis.NotifyPlayer, fmt.Sprintf("%d", uid), myrpc.GetInstance().GetNotifyAddr())
}

func (this *PlayerMgr) DelPlayer(uid uint64) {
	this.players.Delete(uid)
	myredis.GetInstance().HDel(myredis.NotifyPlayer, fmt.Sprintf("%d", uid))
}

func (this *PlayerMgr) Destory() {
	this.players.Range(func(key, value any) bool {
		player := value.(*player_info.Player)
		player.Save()
		return true
	})
}

func (this *PlayerMgr) NotifyAllPlayers(msgid myproto.MsgId, msg myproto.MyMsg) {
	//因为分布式服务器，不止这一个logic节点，所以这边是通知所有gateway，让每个geteway再通知所有玩家
	data, err := msg.Marshal()
	if err != nil {
		mylog.Error(err)
		return
	}
	gateways := myredis.GetInstance().HGetAll("notify_gateway")
	for k, _ := range gateways {
		myrpc.GetInstance().SendMsg(k, 0, msgid, data)
	}
}
