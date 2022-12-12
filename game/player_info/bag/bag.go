package bag

import "gs/proto/myproto"

type Bag struct {
	Unoccupied map[int32]int64
	ItemList   []*myproto.Item
	curId      uint64
}

func NewBag() *Bag {
	return &Bag{
		Unoccupied: make(map[int32]int64),
	}
}
