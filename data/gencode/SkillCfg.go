package gencode

import (
	"encoding/json"
	"gs/define"
	"gs/lib/mylog"
	"io/ioutil"
	"os"
	"sync"
)

type SkillCfg struct {
	SkillSlc []*SkillInfo `json:"Skill"`
	SkillMap map[int32]*SkillInfo
}

type SkillInfo struct {
	ID            int32  // 技能id
	Name          string // 技能名称
	Level         int32  // 技能等级
	BeforeTime    int32  // 前摇时间ms
	AfterTime     int32  // 后摇时间ms
	IsWeaponSkill bool   // 是否为武器平A技能
	CoolDown      int32  // 冷却时间ms
	TargetType    int32  // 目标选择类型
}

var skillCfg *SkillCfg
var skillOnce sync.Once

func GetSkillCfg() *SkillCfg {
	skillOnce.Do(func() {
		skillCfg = new(SkillCfg)
		skillCfg.init()
	})
	return skillCfg
}

func (this *SkillCfg) init() {
	this.SkillMap = make(map[int32]*SkillInfo)
	rootPath := os.Getenv(define.EnvName)
	filePtr, err := os.Open(rootPath + "/data/json/Skill.json")
	if err != nil {
		mylog.Error("load SkillCfg failed", err)
		return
	}
	defer filePtr.Close()
	data, err := ioutil.ReadAll(filePtr)
	if err != nil {
		mylog.Error("load SkillCfg failed", err)
		return
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		mylog.Error("load SkillCfg failed", err)
		return
	}
	for _, v := range this.SkillSlc {
		this.SkillMap[v.ID] = v
	}
}

func (this *SkillCfg) GetSkillById(id int32) (*SkillInfo, bool) {
	if ret, ok := this.SkillMap[id]; ok {
		return ret, ok
	}
	return nil, false
}
