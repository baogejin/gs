package gencode

import (
	"encoding/json"
	"gs/define"
	"gs/lib/mylog"
	"io/ioutil"
	"os"
	"sync"
)

type GlobalCfg struct {
	GlobalSlc []*GlobalInfo `json:"Global"`
	GlobalMap map[string]*GlobalInfo
}

type GlobalInfo struct {
	Key      string    // 键值
	Value    int32     // 数据1
	SlcValue [][]int32 // 数据2，二维数组
	StrValue string    // 字符串数据
}

var globalCfg *GlobalCfg
var globalOnce sync.Once

func GetGlobalCfg() *GlobalCfg {
	globalOnce.Do(func() {
		globalCfg = new(GlobalCfg)
		globalCfg.init()
	})
	return globalCfg
}

func (this *GlobalCfg) init() {
	this.GlobalMap = make(map[string]*GlobalInfo)
	rootPath := os.Getenv(define.EnvName)
	filePtr, err := os.Open(rootPath + "/data/json/Global.json")
	if err != nil {
		mylog.Error("load GlobalCfg failed", err)
		return
	}
	defer filePtr.Close()
	data, err := ioutil.ReadAll(filePtr)
	if err != nil {
		mylog.Error("load GlobalCfg failed", err)
		return
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		mylog.Error("load GlobalCfg failed", err)
		return
	}
	for _, v := range this.GlobalSlc {
		this.GlobalMap[v.Key] = v
	}
}

func (this *GlobalCfg) GetGlobalInfoByKey(key string) (*GlobalInfo, bool) {
	if ret, ok := this.GlobalMap[key]; ok {
		return ret, ok
	}
	return nil, false
}
