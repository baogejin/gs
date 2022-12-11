package battle

import (
	"fmt"
	"gs/data/gencode"
	"gs/lib/mylog"
	"gs/proto/myproto"
	"sort"
	"sync"
	"time"
)

type Battle struct {
	BattleId   uint64
	Units      []*Unit
	ActionList []*BattleAction
	CreateAt   int64
	StartAt    int64
	lock       sync.RWMutex
}

type BattleAction struct {
	ScrUnitId int32
	TarUnitId int32 //单体需要，群技能为0
	SkillId   int32
	TimeStamp int64
}

func (this *Battle) BattleTick() {
	if this.StartAt == 0 {
		return
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	//正在放的技能判断生效
	this.checkBattleAction()
	//单位新释放技能
	this.checkUnitSkill()
	//其他战斗单位状态处理，如自动回血回蓝等
	this.checkUnitStatus()
	//信息打包发送战斗内所有玩家
	//todo
}

func (this *Battle) checkBattleAction() {
	sort.Slice(this.ActionList, func(i, j int) bool {
		return this.ActionList[i].TimeStamp < this.ActionList[j].TimeStamp
	})
	changed := false
	now := time.Now().UnixMilli()
	idx := -1
	for i, v := range this.ActionList {
		if v.TimeStamp > now {
			break
		}
		idx = i
		skillCfg, ok := gencode.GetSkillCfg().GetSkillById(v.SkillId)
		if !ok {
			continue
		}
		src := this.getUintById(v.ScrUnitId)
		if src == nil {
			continue
		}
		if src.IsDead() {
			continue
		}
		//技能生效
		targets := []*Unit{}
		if v.TarUnitId > 0 {
			//单体指定技能
			tar := this.getUintById(v.TarUnitId)
			targets = append(targets, tar)
		} else {
			tars := this.getSkillTargetUnits(skillCfg.TargetType, src)
			targets = append(targets, tars...)
		}
		for _, tar := range targets {
			if tar.IsDead() {
				continue
			}
			if skillCfg.Attack > 0 {
				if tar.HP-int64(skillCfg.Attack) < 0 {
					tar.HP = 0
				} else {
					tar.HP -= int64(skillCfg.Attack)
				}
				mylog.Error(src.Name, " 的", skillCfg.Name, "对 ", tar.Name, " 造成了 ", skillCfg.Attack, " 点伤害")
				changed = true
			}
			if skillCfg.Heal > 0 {
				if tar.HP+int64(skillCfg.Heal) > tar.MaxHP {
					tar.HP = tar.MaxHP
				} else {
					tar.HP += int64(skillCfg.Heal)
				}
				mylog.Info(src.Name, " 的", skillCfg.Name, "对 ", tar.Name, " 造成了 ", skillCfg.Heal, " 点治疗")
				changed = true
			}
		}
	}
	if idx >= 0 {
		this.ActionList = this.ActionList[idx+1:]
	}
	if changed {
		status := "当前血量："
		for _, v := range this.Units {
			status += v.Name + "[" + fmt.Sprintf("%d", v.HP) + "/" + fmt.Sprintf("%d", v.MaxHP) + "] "
		}
		mylog.Warning(status)
	}
}

func (this *Battle) checkUnitSkill() {
	for _, v := range this.Units {
		if !v.CanUseSkill() {
			continue
		}
		skillId := int32(0)
		if v.NextSkill > 0 {
			skillId = v.NextSkill
			v.NextSkill = 0
		} else {
			skillId = v.GetWeaponSkill()
		}
		if skillId == 0 {
			continue
		}
		if skillCfg, ok := gencode.GetSkillCfg().GetSkillById(skillId); ok {
			v.NextFreeTime = time.Now().UnixMilli() + int64(skillCfg.BeforeTime) + int64(skillCfg.AfterTime)
			v.SkillUseTime[skillId] = time.Now().UnixMilli()
			target := this.getSkillTargetUnitId(skillCfg.TargetType, v)
			action := &BattleAction{
				ScrUnitId: v.Id,
				TarUnitId: target,
				SkillId:   skillId,
				TimeStamp: time.Now().UnixMilli() + int64(skillCfg.BeforeTime),
			}
			this.ActionList = append(this.ActionList, action)

			tar := this.getUintById(target)
			if tar != nil {
				mylog.Debug(v.Name, " 开始对 ", tar.Name, " 发动", skillCfg.Name)
			}
		}
	}
}

func (this *Battle) checkUnitStatus() {
	//todo
}

//用于玩家主动施放技能
func (this *Battle) SetNextSkill(uid uint64, skillId int32) myproto.ResultCode {
	if uid == 0 {
		//todo ret uid不能为0
	}
	skillCfg, ok := gencode.GetSkillCfg().GetSkillById(skillId)
	if !ok {
		//todo ret skill没找到啊
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	for _, v := range this.Units {
		if v.Uid == uid {
			if v.NextSkill != 0 {
				//todo ret 下一个技能已经设置了
			}
			//检查cd
			lastUseTime := v.SkillUseTime[skillId]
			if lastUseTime > 0 && lastUseTime+int64(skillCfg.CoolDown) > time.Now().UnixMilli() {
				//todo ret cd
			}
			v.NextSkill = skillId
			return myproto.ResultCode_Success
		}
	}
	//todo ret人没找到
	return myproto.ResultCode_PlayerNotFound
}

func (this *Battle) getUintById(unitId int32) *Unit {
	for _, v := range this.Units {
		if v.Id == unitId {
			return v
		}
	}
	return nil
}
