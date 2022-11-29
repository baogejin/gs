package gateway

import (
	"gs/define"
	"gs/lib/mylog"
	"gs/lib/myrpc"
	"net/http"

	"golang.org/x/net/websocket"
)

type GatewayServer struct {
}

func (this *GatewayServer) Init() {
}

func OnNewConn(ws *websocket.Conn) {
	client := &Client{
		ws: ws,
	}
	mylog.Info("new client ", ws.RemoteAddr())
	client.Start()
}

func (this *GatewayServer) Run() {
	mylog.Info("gateway server run")
	myrpc.GetInstance().RegisterClient(define.NodeLogic, nil, nil)
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
