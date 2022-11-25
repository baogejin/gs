package logic

import (
	"fmt"
)

type LogicServer struct {
}

func (this *LogicServer) Init() {

}

func (this *LogicServer) Run() {
	fmt.Println("logic server run")
}

func (this *LogicServer) Destory() {

}
