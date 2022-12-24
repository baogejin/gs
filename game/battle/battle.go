package battle

import (
	"fmt"
	"gs/data/gencode"
	"gs/define"
	"gs/lib/mylog"
	"gs/lib/myrpc"
	"gs/proto/myproto"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

var curBattleId uint64 = 0

type Battle struct {
	BattleId   uint64
	LevelId    int32
	curUnitId  int32
	Units      []*Unit
	ActionList []*BattleAction
	CreateAt   int64
	StartAt    int64
	End        bool
	lock       sync.RWMutex
}

func CreatePveBattle(levelId int32, playerUnits ...*Unit) *Battle {
	if len(playerUnits) == 0 {
		return nil
	}
	battleId := atomic.AddUint64(&curBattleId, 1)
	b := &Battle{
		BattleId:  battleId,
		LevelId:   levelId,
		curUnitId: 0,
		CreateAt:  time.Now().UnixMilli(),
	}
	for _, unit := range playerUnits {
		unit.Id = b.genUnitId()
		unit.Team = 0
		unit.SkillUseTime = make(map[int32]int64)
	}
	b.Units = append(b.Units, playerUnits...)

	//todo 根据关卡配置增加怪物
	b.Units = append(b.Units, &Unit{
		Id:           b.genUnitId(),
		Name:         "魔王",
		Team:         1,
		Position:     4,
		UnitType:     myproto.UnitType_UnitMonster,
		WeaponSkill:  3,
		SkillUseTime: make(map[int32]int64),
		HP:           100,
		MaxHP:        100,
	})

	return b
}

//todo CreatePvpBattle

func (this *Battle) BattleInfoNotify() {
	push := &myproto.BattleInfoPUSH{
		BattleId: this.BattleId,
		LevelId:  this.LevelId,
	}
	for _, v := range this.Units {
		push.Units = append(push.Units, v.Proto())
	}
	data, err := push.Marshal()
	if err != nil {
		mylog.Error(err)
		return
	}
	this.notifyAllUnits(myproto.MsgId_Msg_BattleInfoPUSH, data)
}

func (this *Battle) notifyAllUnits(msgId myproto.MsgId, data []byte) {
	for _, v := range this.Units {
		if v.NotifyAddr != "" && v.Uid != 0 {
			myrpc.GetInstance().SendMsg(v.NotifyAddr, v.Uid, msgId, define.NodeGateway, data)
		}
	}
}

func (this *Battle) genUnitId() int32 {
	unitId := atomic.AddInt32(&this.curUnitId, 1)
	return unitId
}

func (this *Battle) Start() {
	if this.StartAt > 0 {
		return
	}
	this.StartAt = time.Now().UnixMilli()
	push := &myproto.BattleStartPUSH{}
	data, err := push.Marshal()
	if err != nil {
		mylog.Error(err)
		return
	}
	this.notifyAllUnits(myproto.MsgId_Msg_BattleStartPUSH, data)
}

type BattleAction struct {
	Src       int32
	Tar       []int32
	SkillId   int32
	TimeStamp int64
}

func (this *Battle) BattleTick() {
	if this.StartAt == 0 {
		return
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	//结束检查
	if this.checkEnd() {
		return
	}
	//正在放的技能判断生效
	effects := this.checkBattleAction()
	//单位新释放技能
	skills := this.checkUnitSkill()
	//其他战斗单位状态处理，如自动回血回蓝等
	this.checkUnitStatus()
	//信息打包发送战斗内所有玩家
	if len(effects) > 0 || len(skills) > 0 {
		push := &myproto.BattleActionPUSH{
			BattleId: this.BattleId,
			Effects:  effects,
			Skills:   skills,
		}
		data, err := push.Marshal()
		if err != nil {
			mylog.Error(err)
		} else {
			this.notifyAllUnits(myproto.MsgId_Msg_BattleActionPUSH, data)
		}
	}
}

func (this *Battle) checkEnd() bool {
	if this.End {
		return true
	}
	teamMap := make(map[int32]bool)
	for _, v := range this.Units {
		if !v.IsDead() {
			teamMap[v.Team] = true
		}
	}
	if len(teamMap) < 2 {
		this.End = true
		winTeam := 1
		if teamMap[0] {
			winTeam = 0
		}
		winPush := &myproto.BattleFinishPUSH{
			BattleId: this.BattleId,
			Win:      true,
		}
		winData, err := winPush.Marshal()
		if err != nil {
			mylog.Error(err)
		} else {
			for _, v := range this.Units {
				if v.Uid > 0 && v.NotifyAddr != "" && v.Team == int32(winTeam) {
					myrpc.GetInstance().SendMsg(v.NotifyAddr, v.Uid, myproto.MsgId_Msg_BattleFinishPUSH, define.NodeGateway, winData)
				}
			}
		}
		losePush := &myproto.BattleFinishPUSH{
			BattleId: this.BattleId,
			Win:      false,
		}
		loseData, err := losePush.Marshal()
		if err != nil {
			mylog.Error(err)
		} else {
			for _, v := range this.Units {
				if v.Uid > 0 && v.NotifyAddr != "" && v.Team != int32(winTeam) {
					myrpc.GetInstance().SendMsg(v.NotifyAddr, v.Uid, myproto.MsgId_Msg_BattleFinishPUSH, define.NodeGateway, loseData)
				}
			}
		}
	}
	return this.End
}

func (this *Battle) checkBattleAction() []*myproto.BattleSkillEffect {
	effects := []*myproto.BattleSkillEffect{}
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
		src := this.getUintById(v.Src)
		if src == nil {
			continue
		}
		if src.IsDead() {
			continue
		}
		//技能生效
		targets := []*Unit{}
		for _, unitId := range v.Tar {
			unit := this.getUintById(unitId)
			if unit != nil {
				targets = append(targets, unit)
			}
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
				effect := &myproto.BattleSkillEffect{
					Src:     v.Src,
					SkillId: v.SkillId,
					Tar:     tar.Id,
					IsHeal:  false,
					Value:   int64(skillCfg.Attack),
				}
				effects = append(effects, effect)
			}
			if skillCfg.Heal > 0 {
				if tar.HP+int64(skillCfg.Heal) > tar.MaxHP {
					tar.HP = tar.MaxHP
				} else {
					tar.HP += int64(skillCfg.Heal)
				}
				mylog.Info(src.Name, " 的", skillCfg.Name, "对 ", tar.Name, " 造成了 ", skillCfg.Heal, " 点治疗")
				changed = true
				effect := &myproto.BattleSkillEffect{
					Src:     v.Src,
					SkillId: v.SkillId,
					Tar:     tar.Id,
					IsHeal:  true,
					Value:   int64(skillCfg.Heal),
				}
				effects = append(effects, effect)
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
	return effects
}

func (this *Battle) checkUnitSkill() []*myproto.BattleSkillStart {
	skills := []*myproto.BattleSkillStart{}
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
				Src:       v.Id,
				Tar:       target,
				SkillId:   skillId,
				TimeStamp: time.Now().UnixMilli() + int64(skillCfg.BeforeTime),
			}
			skills = append(skills, &myproto.BattleSkillStart{
				Src:     v.Id,
				SkillId: skillId,
				Tar:     target,
			})
			this.ActionList = append(this.ActionList, action)
			for _, t := range target {
				tar := this.getUintById(t)
				if tar != nil {
					mylog.Debug(v.Name, " 开始对 ", tar.Name, " 发动", skillCfg.Name)
				}
			}
		}
	}
	return skills
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
