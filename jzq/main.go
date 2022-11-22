package main

import (
	"fmt"
	"gs/data/gencode"
)

func main() {
	if e, ok := gencode.GetExampleCfg().GetExampleById(1); ok {
		fmt.Println(e.Name)
	}
	if a, ok := gencode.GetAnotherExampleCfg().GetAnotherById(1); ok {
		fmt.Println(a.Name, a.Age)
	}
	if g, ok := gencode.GetGlobalCfg().GetGlobalInfoByKey("TestKey"); ok {
		fmt.Println(g.Value, g.Value2)
	}
}
