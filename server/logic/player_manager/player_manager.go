package player_manager

import (
	"gs/game/player_info"
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
}

func (this *PlayerMgr) DelPlayer(uid uint64) {
	this.players.Delete(uid)
}

func (this *PlayerMgr) Destory() {
	this.players.Range(func(key, value any) bool {
		player := value.(*player_info.Player)
		player.Save()
		return true
	})
}

func (this *PlayerMgr) Broadcast(msgid myproto.MsgId, msg myproto.MyMsg) {
	this.players.Range(func(key, value any) bool {
		player := value.(*player_info.Player)
		player.SendMsg(msgid, msg)
		return true
	})
}
