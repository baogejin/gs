package player_info

import "gs/proto/myproto"

func (this *Player) AddItems(items ...*myproto.Item) {
	changeItems := this.Bag.AddItems(items...)
	//通知前端
	if len(changeItems) > 0 {
		this.SendMsg(myproto.MsgId_Msg_ItemUpdatePUSH, &myproto.ItemUpdatePUSH{
			UpdateType: myproto.ItemUpdateType_ItemAdd,
			Items:      changeItems,
		})
	}
}

func (this *Player) CostItems(items ...*myproto.Item) bool {
	changeItems, ok := this.Bag.CostItems(items...)
	if ok && len(changeItems) > 0 {
		//通知前端
		this.SendMsg(myproto.MsgId_Msg_ItemUpdatePUSH, &myproto.ItemUpdatePUSH{
			UpdateType: myproto.ItemUpdateType_ItemDel,
			Items:      changeItems,
		})
	}
	return ok
}
