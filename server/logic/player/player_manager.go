package player

import "sync"

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
