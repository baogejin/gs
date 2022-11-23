package main

import (
	"gs/data/gencode"
	"gs/lib/mylog"
)

func main() {

	if c, ok := gencode.GetGlobalCfg().GetGlobalInfoByKey(gencode.TestKey2); ok {
		mylog.Error(c.StrValue)
	}
}
