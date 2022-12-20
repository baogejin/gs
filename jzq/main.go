package main

import (
	"fmt"
	"gs/game/battle"
)

func main() {
	b := battle.CreateBattle()
	fmt.Println(b.BattleId)
	b1 := battle.CreateBattle()
	fmt.Println(b1.BattleId)
	// b := &battle.Battle{
	// 	CreateAt: time.Now().UnixMilli(),
	// 	StartAt:  time.Now().UnixMilli(),
	// }
	// b.Units = append(b.Units, &battle.Unit{
	// 	Id:           1,
	// 	Name:         "魔王",
	// 	Team:         0,
	// 	Position:     3,
	// 	WeaponSkill:  3,
	// 	SkillUseTime: make(map[int32]int64),
	// 	HP:           1000,
	// 	MaxHP:        1000,
	// })
	// b.Units = append(b.Units, &battle.Unit{
	// 	Id:           2,
	// 	Name:         "恶魔",
	// 	Team:         0,
	// 	Position:     1,
	// 	WeaponSkill:  1,
	// 	SkillUseTime: make(map[int32]int64),
	// 	HP:           500,
	// 	MaxHP:        500,
	// })
	// b.Units = append(b.Units, &battle.Unit{
	// 	Id:           3,
	// 	Name:         "勇者",
	// 	Team:         1,
	// 	Position:     1,
	// 	WeaponSkill:  1,
	// 	SkillUseTime: make(map[int32]int64),
	// 	HP:           200,
	// 	MaxHP:        200,
	// })
	// b.Units = append(b.Units, &battle.Unit{
	// 	Id:           4,
	// 	Name:         "射手",
	// 	Team:         1,
	// 	Position:     3,
	// 	WeaponSkill:  2,
	// 	SkillUseTime: make(map[int32]int64),
	// 	HP:           200,
	// 	MaxHP:        200,
	// })
	// b.Units = append(b.Units, &battle.Unit{
	// 	Id:           5,
	// 	Name:         "牧师",
	// 	Team:         1,
	// 	Position:     5,
	// 	WeaponSkill:  4,
	// 	SkillUseTime: make(map[int32]int64),
	// 	HP:           200,
	// 	MaxHP:        200,
	// })
	// myticker.GetInstance().AddTicker(time.Millisecond*50, b.BattleTick)
	// time.Sleep(time.Hour)
}
