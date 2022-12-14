package player_info

import "gs/proto/myproto"

func (this *Player) AddItems(items ...*myproto.Item) {
	this.Bag.AddItems(items...)
	//todo 通知前端
}

func (this *Player) CostItems(items ...*myproto.Item) bool {
	_, ok := this.Bag.CostItems(items...)
	if ok {
		//todo 通知前端
	}
	return ok
}
