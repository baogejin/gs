package bag

import (
	"gs/data/gencode"
	"gs/lib/mylog"
	"gs/proto/myproto"
	"sync/atomic"
)

type Bag struct {
	Unoccupied   map[int32]int64
	StackItems   map[int32]*myproto.Item
	UnstackItems map[uint64]*myproto.Item
	curId        uint64
	//todo 格子上限
}

func NewBag() *Bag {
	return &Bag{
		Unoccupied:   make(map[int32]int64),
		StackItems:   make(map[int32]*myproto.Item),
		UnstackItems: make(map[uint64]*myproto.Item),
	}
}

//todo 背包上限，溢出走邮件逻辑
func (this *Bag) AddItems(items ...*myproto.Item) []*myproto.Item {
	stack := make(map[int32]int64)
	ret := make([]*myproto.Item, 0)
	for _, v := range items {
		itemCfg, ok := gencode.GetItemCfg().GetItemById(v.ItemId)
		if !ok {
			mylog.Warning("item not found ", v.ItemId)
			continue
		}
		switch myproto.ItemType(itemCfg.Type) {
		case myproto.ItemType_UnoccupiedItem:
			this.Unoccupied[v.ItemId] += v.Num
			stack[v.ItemId] += v.Num
		case myproto.ItemType_StackItem:
			item, ok := this.StackItems[v.ItemId]
			if ok {
				item.Num += v.Num
			} else {
				item = &myproto.Item{ItemId: v.ItemId, Num: v.Num}
				this.StackItems[v.ItemId] = item
			}
			stack[v.ItemId] += v.Num
		case myproto.ItemType_UnstackItem, myproto.ItemType_EquipItem:
			for i := 0; i < int(v.Num); i++ {
				id := atomic.AddUint64(&this.curId, 1)
				item := &myproto.Item{Id: id, ItemId: v.ItemId, Num: v.Num}
				this.UnstackItems[item.Id] = item
				ret = append(ret, item)
			}
		default:
			mylog.Warning("item type undefined ", v.ItemId)
		}
	}
	for itemId, num := range stack {
		ret = append(ret, &myproto.Item{ItemId: itemId, Num: num})
	}
	return ret
}

func (this *Bag) CostItems(items ...*myproto.Item) ([]*myproto.Item, bool) {
	unoccupied, stack, unstack, ok := checkCostItemStack(items...)
	if !ok {
		return nil, false
	}
	ret := make([]*myproto.Item, 0)
	//先检查够不够
	for itemId, num := range unoccupied {
		if this.Unoccupied[itemId] < num {
			return nil, false
		}
	}
	for itemId, num := range stack {
		if item, ok := this.StackItems[itemId]; ok {
			if item.Num < num {
				return nil, false
			}
		} else {
			return nil, false
		}
	}
	for _, v := range unstack {
		if item, ok := this.UnstackItems[v.Id]; ok {
			if item.ItemId != v.ItemId {
				return nil, false
			}
		} else {
			return nil, false
		}
	}

	//再实际扣东西
	for itemId, num := range unoccupied {
		this.Unoccupied[itemId] -= num
		ret = append(ret, &myproto.Item{ItemId: itemId, Num: num})
	}
	for itemId, num := range stack {
		if item, ok := this.StackItems[itemId]; ok {
			item.Num -= num
			if item.Num == 0 {
				delete(this.StackItems, itemId)
			}
			ret = append(ret, &myproto.Item{ItemId: itemId, Num: num})
		}
	}
	for _, v := range unstack {
		delete(this.UnstackItems, v.Id)
		ret = append(ret, v)
	}
	return ret, true
}

func (this *Bag) GetSpaceUsed() int32 {
	return int32(len(this.StackItems) + len(this.UnstackItems))
}

func checkCostItemStack(items ...*myproto.Item) (map[int32]int64, map[int32]int64, []*myproto.Item, bool) {
	unstackItems := make([]*myproto.Item, 0)
	unoccupiedMap := make(map[int32]int64)
	stackMap := make(map[int32]int64)
	unstackRepeatCheck := make(map[uint64]bool)
	for _, v := range items {
		itemCfg, ok := gencode.GetItemCfg().GetItemById(v.ItemId)
		if !ok {
			mylog.Warning("item not found ", v.ItemId)
			return nil, nil, nil, false
		}
		if myproto.ItemType(itemCfg.Type) == myproto.ItemType_UnoccupiedItem {
			unoccupiedMap[v.ItemId] += v.Num
		} else if myproto.ItemType(itemCfg.Type) == myproto.ItemType_StackItem {
			stackMap[v.ItemId] += v.Num
		} else {
			if v.Id == 0 {
				mylog.Warning("unstack item id zero")
				return nil, nil, nil, false
			}
			if v.Num != 1 {
				mylog.Warning("unstack item num not one")
				return nil, nil, nil, false
			}
			if unstackRepeatCheck[v.Id] {
				mylog.Warning("unstack item id repeat,id ", v.Id)
				return nil, nil, nil, false
			}
			unstackRepeatCheck[v.Id] = true
			unstackItems = append(unstackItems, v)
		}
	}
	return unoccupiedMap, stackMap, unstackItems, true
}
