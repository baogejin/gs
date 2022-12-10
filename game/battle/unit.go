package battle

import (
	"gs/data/gencode"
	"gs/proto/myproto"
	"time"
)

type Unit struct {
	Id           int32
	Uid          uint64
	Name         string
	Team         int32 //队伍，区分敌我
	Position     int32 //位置
	UnitType     myproto.UnitType
	NextSkill    int32
	SkillUseTime map[int32]int64 //技能上一次使用时间，计算cd用
	NextFreeTime int64           //可以施放下一个技能的时间
}

func (this *Unit) CanUseSkill() bool {
	//todo 如果死亡
	if this.NextFreeTime > time.Now().UnixMilli() {
		return false
	}
	return true
}

func (this *Unit) GetWeaponSkill() int32 {
	skillId := int32(1) //todo 后续根据武器获得武器技能id
	if skillCfg, ok := gencode.GetSkillCfg().GetSkillById(skillId); ok {
		lastUseTime := this.SkillUseTime[skillId]
		if time.Now().UnixMilli() >= lastUseTime+int64(skillCfg.CoolDown) {
			return skillId
		} else {
			return 0
		}
	}
	return 0
}

func (this *Unit) IsFront() bool {
	return this.Position < 3
}

func (this *Unit) GetTargetWeight() uint {
	return 100 //todo 后续根据位置返回成为目标的权重
}
