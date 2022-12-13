package gencode

import (
	"encoding/json"
	"gs/define"
	"gs/lib/mylog"
	"io/ioutil"
	"os"
	"sync"
)

type ItemCfg struct {
	ItemSlc []*ItemInfo `json:"Item"`
	ItemMap map[int32]*ItemInfo
}

type ItemInfo struct {
	ID      int32  // 物品id
	Name    string // 物品名字
	Type    int32  // 物品类型
	Quality int32  // 品质
}

var itemCfg *ItemCfg
var itemOnce sync.Once

func GetItemCfg() *ItemCfg {
	itemOnce.Do(func() {
		itemCfg = new(ItemCfg)
		itemCfg.init()
	})
	return itemCfg
}

func (this *ItemCfg) init() {
	this.ItemMap = make(map[int32]*ItemInfo)
	rootPath := os.Getenv(define.EnvName)
	filePtr, err := os.Open(rootPath + "/data/json/Item.json")
	if err != nil {
		mylog.Error("load ItemCfg failed", err)
		return
	}
	defer filePtr.Close()
	data, err := ioutil.ReadAll(filePtr)
	if err != nil {
		mylog.Error("load ItemCfg failed", err)
		return
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		mylog.Error("load ItemCfg failed", err)
		return
	}
	for _, v := range this.ItemSlc {
		this.ItemMap[v.ID] = v
	}
}

func (this *ItemCfg) GetItemById(id int32) (*ItemInfo, bool) {
	if ret, ok := this.ItemMap[id]; ok {
		return ret, ok
	}
	return nil, false
}
