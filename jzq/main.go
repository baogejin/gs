package main

import (
	"gs/data/gencode"
	"gs/lib/mylog"
)

func main() {

	if c, ok := gencode.GetGlobalCfg().GetGlobalInfoByKey(gencode.TestKey2); ok {
		mylog.Debug(c.StrValue)
		mylog.Info(c.StrValue)
		mylog.Notice(c.StrValue)
		mylog.Warning(c.StrValue)
		mylog.Error(c.StrValue)
		mylog.Alert(c.StrValue)
		mylog.Critical(c.StrValue)
		mylog.Emergency(c.StrValue)
	}
}
