package main

import (
	"fmt"
	"gs/proto/myproto"
	"gs/server/gateway"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://localhost/"
	url := "ws://localhost:12345/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	req := &myproto.RegisterREQ{Account: "jzq", Password: "123"}
	data, _ := req.Marshal()
	msgByte := gateway.PackMsg(uint32(myproto.MsgId_Msg_RegisterREQ), data)
	if _, err := ws.Write(msgByte); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 1)
	var buf = make([]byte, 512)
	var n int
	if n, err = ws.Read(buf); err != nil {
		log.Fatal(err)
	}
	if n > 0 {
		pack := gateway.UnpackMsg(buf[4:])
		ack := &myproto.RegisterACK{}
		ack.Unmarshal(pack.Data)
		fmt.Println(ack.Ret)
	}

	req1 := &myproto.LoginREQ{Account: "jzq", Password: "123"}
	data, _ = req1.Marshal()
	msgByte = gateway.PackMsg(uint32(myproto.MsgId_Msg_LoginREQ), data)
	if _, err := ws.Write(msgByte); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 1)
	if n, err = ws.Read(buf); err != nil {
		log.Fatal(err)
	}
	if n > 0 {
		pack := gateway.UnpackMsg(buf[4:])
		ack := &myproto.LoginACK{}
		ack.Unmarshal(pack.Data)
		fmt.Println(ack)
	}

	time.Sleep(time.Second * 10)
	if n, err = ws.Read(buf); err != nil {
		log.Fatal(err)
	}
	if n > 0 {
		pack := gateway.UnpackMsg(buf[4:])
		ack := &myproto.KickPUSH{}
		ack.Unmarshal(pack.Data)
		fmt.Println(ack)
	}

	// req2 := &myproto.CreateRoleREQ{Name: "jzq"}
	// data, _ := req2.Marshal()
	// msgByte := gateway.PackMsg(uint32(myproto.MsgId_Msg_CreateRoleREQ), data)
	// if _, err := ws.Write(msgByte); err != nil {
	// 	log.Fatal(err)
	// }
	// time.Sleep(time.Second * 1)
	// if n, err = ws.Read(buf); err != nil {
	// 	log.Fatal(err)
	// }
	// if n > 0 {
	// 	pack := gateway.UnpackMsg(buf[4:])
	// 	ack := &myproto.CreateRoleACK{}
	// 	ack.Unmarshal(pack.Data)
	// 	fmt.Println(ack)
	// }

	ws.Close()

}
