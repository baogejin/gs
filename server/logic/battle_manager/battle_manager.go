package battle_manager

import (
	"gs/game/battle"
	"gs/lib/myticker"
	"sync"
	"time"
)

type BattleMgr struct {
	battles sync.Map
}

var battleMgr *BattleMgr
var once sync.Once

func GetMgr() *BattleMgr {
	once.Do(func() {
		battleMgr = new(BattleMgr)
		battleMgr.init()
	})
	return battleMgr
}

func (this *BattleMgr) init() {
	myticker.GetInstance().AddTicker(time.Millisecond*50, this.RunTick)
}

func (this *BattleMgr) RunTick() {
	endBattles := make([]uint64, 0)
	this.battles.Range(func(key, value any) bool {
		b := value.(*battle.Battle)
		if b.End {
			endBattles = append(endBattles, b.BattleId)
			return true
		}
		b.BattleTick()
		return true
	})
	for _, id := range endBattles {
		this.battles.Delete(id)
	}
}

func (this *BattleMgr) GetBattle(battleId uint64) *battle.Battle {
	b, ok := this.battles.Load(battleId)
	if !ok {
		return nil
	}
	return b.(*battle.Battle)
}

func (this *BattleMgr) AddBattle(b *battle.Battle) {
	this.battles.Store(b.BattleId, b)
}
