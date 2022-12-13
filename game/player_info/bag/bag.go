package bag

import "gs/proto/myproto"

type Bag struct {
	Unoccupied map[int32]int64
	ItemList   []*myproto.Item
	curId      uint64
	//todo 格子上限
}

func NewBag() *Bag {
	return &Bag{
		Unoccupied: make(map[int32]int64),
	}
}

func (this *Bag) AddItems(items ...myproto.Item) {

}

func (this *Bag) CostItems(items ...myproto.Item) bool {
	return true
}
