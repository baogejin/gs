package myconfig

import (
	"encoding/json"
	"fmt"
	"gs/define"
	"io/ioutil"
	"os"
	"sync"
)

type ConfigInfo struct {
	LogLevel     string
	RedisAddress string
	RpcPortStart int32
}

var cfg *ConfigInfo
var once sync.Once

func Get() *ConfigInfo {
	once.Do(func() {
		cfg = new(ConfigInfo)
		cfg.init()
	})
	return cfg
}

func (this *ConfigInfo) init() {
	rootPath := os.Getenv(define.EnvName)
	filePtr, err := os.Open(rootPath + "/config/config.json")
	if err != nil {
		fmt.Println("config init failed", err)
		return
	}
	defer filePtr.Close()
	data, err := ioutil.ReadAll(filePtr)
	if err != nil {
		fmt.Println("config init failed", err)
		return
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("config init failed", err)
		return
	}
}
