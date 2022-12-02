package main

import (
	"encoding/json"
	"fmt"
	"gs/server/logic/player"
	"time"
)

func main() {
	player := &player.Player{
		Uid:      1,
		Name:     "dasdasd",
		CreateAt: time.Now().Unix(),
	}
	jsonData, _ := json.Marshal(player)
	fmt.Println(string(jsonData))
}
