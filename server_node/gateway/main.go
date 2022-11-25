package gateway

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type GatewayServer struct {
}

func (this *GatewayServer) Init() {
}

func OnNewConn(ws *websocket.Conn) {
	client := &ClientConn{
		ws: ws,
	}
	client.Start()
}

func (this *GatewayServer) Run() {
	fmt.Println("gateway server run")
	//对外websocket
	go func() {
		http.Handle("/", websocket.Handler(OnNewConn))
		err := http.ListenAndServe(":12345", nil)
		if err != nil {
			panic("ListenAndServer: " + err.Error())
		}
	}()
}

func (this *GatewayServer) Destory() {

}
