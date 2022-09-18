package main

import (
	"gs/serverNode/gateway"
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
	msgByte := gateway.PackMsg(0, []byte("hello world"))
	if _, err := ws.Write(msgByte); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 10)
	if _, err := ws.Write(msgByte); err != nil {
		log.Fatal(err)
	}
	ws.Close()

	// var msg = make([]byte, 512)
	// var n int
	// if n, err = ws.Read(msg); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Received: %s.\n", msg[:n])
}
