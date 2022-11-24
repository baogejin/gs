package gencode

import (
	"encoding/json"
	"gs/define"
	"gs/lib/mylog"
	"io/ioutil"
	"os"
	"sync"
)

type AnotherExampleCfg struct {
	AnotherSlc []*AnotherInfo `json:"Another"`
	AnotherMap map[int32]*AnotherInfo
}

type AnotherInfo struct {
	ID   int32  // ID
	Name string // 姓名
	Age  int32  // 年龄
}

var anotherexampleCfg *AnotherExampleCfg
var anotherexampleOnce sync.Once

func GetAnotherExampleCfg() *AnotherExampleCfg {
	anotherexampleOnce.Do(func() {
		anotherexampleCfg = new(AnotherExampleCfg)
		anotherexampleCfg.init()
	})
	return anotherexampleCfg
}

func (this *AnotherExampleCfg) init() {
	this.AnotherMap = make(map[int32]*AnotherInfo)
	rootPath := os.Getenv(define.EnvName)
	filePtr, err := os.Open(rootPath + "/data/json/AnotherExample.json")
	if err != nil {
		mylog.Error("load AnotherExampleCfg failed", err)
		return
	}
	defer filePtr.Close()
	data, err := ioutil.ReadAll(filePtr)
	if err != nil {
		mylog.Error("load AnotherExampleCfg failed", err)
		return
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		mylog.Error("load AnotherExampleCfg failed", err)
		return
	}
	for _, v := range this.AnotherSlc {
		this.AnotherMap[v.ID] = v
	}
}

func (this *AnotherExampleCfg) GetAnotherById(id int32) (*AnotherInfo, bool) {
	if ret, ok := this.AnotherMap[id]; ok {
		return ret, ok
	}
	return nil, false
}
