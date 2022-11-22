package gencode

import (
	"encoding/json"
	"gs/define"
	"io/ioutil"
	"os"
	"sync"
)

type ExampleCfg struct {
	ExampleSlc []*ExampleInfo `json:"Example"`
	ExampleMap map[int32]*ExampleInfo
	AbbSlc     []*AbbInfo `json:"Abb"`
	AbbMap     map[int32]*AbbInfo
}

type ExampleInfo struct {
	ID         int32            // 数字
	Name       string           // 字符串
	Slc1       []int32          // 数组
	Slc2       []float32        // 数组
	DoubleSlc1 [][]int32        // 二维数组
	DoubleSlc2 [][]string       // 二维数组
	IsBool     bool             // 布尔
	Map1       map[int32]int32  // map类型
	Map2       map[int32]string // map类型
}

type AbbInfo struct {
	ID  int32  // ID
	Sth string // 参数1
}

var exampleCfg *ExampleCfg
var exampleOnce sync.Once

func GetExampleCfg() *ExampleCfg {
	exampleOnce.Do(func() {
		exampleCfg = new(ExampleCfg)
		exampleCfg.init()
	})
	return exampleCfg
}

func (this *ExampleCfg) init() {
	this.ExampleMap = make(map[int32]*ExampleInfo)
	this.AbbMap = make(map[int32]*AbbInfo)
	rootPath := os.Getenv(define.EnvName)
	filePtr, err := os.Open(rootPath + "/data/json/Example.json")
	if err != nil {
		panic(err)
	}
	defer filePtr.Close()
	data, err := ioutil.ReadAll(filePtr)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		panic(err)
	}
	for _, v := range this.ExampleSlc {
		this.ExampleMap[v.ID] = v
	}
	for _, v := range this.AbbSlc {
		this.AbbMap[v.ID] = v
	}
}

func (this *ExampleCfg) GetExampleById(id int32) (*ExampleInfo, bool) {
	if ret, ok := this.ExampleMap[id]; ok {
		return ret, ok
	}
	return nil, false
}

func (this *ExampleCfg) GetAbbById(id int32) (*AbbInfo, bool) {
	if ret, ok := this.AbbMap[id]; ok {
		return ret, ok
	}
	return nil, false
}
