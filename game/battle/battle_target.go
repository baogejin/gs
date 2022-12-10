package battle

import (
	"gs/lib/mychooser"
	"gs/lib/mylog"
	"gs/proto/myproto"
)

//添加action列表时提前计算指定的单体技能目标的unitid，群体技能返回0
func (this *Battle) getSkillTargetUnitId(targetType int32, src *Unit) int32 {
	switch myproto.TargetType(targetType) {
	case myproto.TargetType_TargetSelf:
		return src.Id
	case myproto.TargetType_EnemySingle, myproto.TargetType_EnemySingleFront, myproto.TargetType_EnemySingleBehind,
		myproto.TargetType_AllySingle, myproto.TargetType_AllySingleFront, myproto.TargetType_AllySingleBehind:
		units := this.getSkillTargetUnits(targetType, src)
		if len(units) == 1 {
			return units[0].Id
		}
	}
	return 0
}

func (this *Battle) getSkillTargetUnits(targetType int32, src *Unit) []*Unit {
	switch myproto.TargetType(targetType) {
	case myproto.TargetType_EnemySingle:
		unit := this.GetEnemySingle(src.Team)
		if unit != nil {
			return []*Unit{unit}
		}
	case myproto.TargetType_EnemySingleFront:
		unit := this.GetEnemySingleFront(src.Team)
		if unit != nil {
			return []*Unit{unit}
		}
	case myproto.TargetType_EnemySingleBehind:
		unit := this.GetEnemySingleBehind(src.Team)
		if unit != nil {
			return []*Unit{unit}
		}
	case myproto.TargetType_EnemyAll:
		front, behind := this.GetEnemy(src.Team)
		return append(front, behind...)
	case myproto.TargetType_EnemyFrontAll:
		front, _ := this.GetEnemy(src.Team)
		return front
	case myproto.TargetType_EnemyBehindAll:
		_, behind := this.GetEnemy(src.Team)
		return behind
	case myproto.TargetType_TargetSelf:
		return []*Unit{src}
	case myproto.TargetType_AllySingle:
		unit := this.GetAllySingle(src.Team)
		if unit != nil {
			return []*Unit{unit}
		}
	case myproto.TargetType_AllySingleFront:
		unit := this.GetAllySingleFront(src.Team)
		if unit != nil {
			return []*Unit{unit}
		}
	case myproto.TargetType_AllySingleBehind:
		unit := this.GetAllySingleBehind(src.Team)
		if unit != nil {
			return []*Unit{unit}
		}
	case myproto.TargetType_AllyAll:
		front, behind := this.GetAlly(src.Team)
		return append(front, behind...)
	case myproto.TargetType_AllyFrontAll:
		front, _ := this.GetAlly(src.Team)
		return front
	case myproto.TargetType_AllyBehindAll:
		_, behind := this.GetAlly(src.Team)
		return behind
	}
	return nil
}

func (this *Battle) GetEnemy(team int32) ([]*Unit, []*Unit) {
	front := make([]*Unit, 0)
	behind := make([]*Unit, 0)
	for _, v := range this.Units {
		if v.Team == team {
			continue
		}
		if v.IsFront() {
			front = append(front, v)
		} else {
			behind = append(behind, v)
		}
	}
	return front, behind
}

func (this *Battle) GetAlly(team int32) ([]*Unit, []*Unit) {
	front := make([]*Unit, 0)
	behind := make([]*Unit, 0)
	for _, v := range this.Units {
		if v.Team != team {
			continue
		}
		if v.IsFront() {
			front = append(front, v)
		} else {
			behind = append(behind, v)
		}
	}
	return front, behind
}

func (this *Battle) GetEnemySingle(team int32) *Unit {
	front, behind := this.GetEnemy(team)
	targets := append(front, behind...)
	if len(targets) == 0 {
		return nil
	}
	c := &mychooser.MyChooser{}
	for _, v := range targets {
		c.Add(v, v.GetTargetWeight())
	}
	pick, err := c.Pick()
	if err != nil {
		mylog.Error(err)
	}
	return pick.(*Unit)
}

func (this *Battle) GetEnemySingleFront(team int32) *Unit {
	front, behind := this.GetEnemy(team)
	targets := front
	if len(targets) == 0 {
		targets = behind
	}
	if len(targets) == 0 {
		return nil
	}
	c := &mychooser.MyChooser{}
	for _, v := range targets {
		c.Add(v, v.GetTargetWeight())
	}
	pick, err := c.Pick()
	if err != nil {
		mylog.Error(err)
	}
	return pick.(*Unit)
}

func (this *Battle) GetEnemySingleBehind(team int32) *Unit {
	front, behind := this.GetEnemy(team)
	targets := behind
	if len(targets) == 0 {
		targets = front
	}
	if len(targets) == 0 {
		return nil
	}
	c := &mychooser.MyChooser{}
	for _, v := range targets {
		c.Add(v, v.GetTargetWeight())
	}
	pick, err := c.Pick()
	if err != nil {
		mylog.Error(err)
	}
	return pick.(*Unit)
}

func (this *Battle) GetAllySingle(team int32) *Unit {
	front, behind := this.GetAlly(team)
	targets := append(front, behind...)
	if len(targets) == 0 {
		return nil
	}
	c := &mychooser.MyChooser{}
	for _, v := range targets {
		c.Add(v, v.GetTargetWeight())
	}
	pick, err := c.Pick()
	if err != nil {
		mylog.Error(err)
	}
	return pick.(*Unit)
}

func (this *Battle) GetAllySingleFront(team int32) *Unit {
	front, behind := this.GetAlly(team)
	targets := front
	if len(targets) == 0 {
		targets = behind
	}
	if len(targets) == 0 {
		return nil
	}
	c := &mychooser.MyChooser{}
	for _, v := range targets {
		c.Add(v, v.GetTargetWeight())
	}
	pick, err := c.Pick()
	if err != nil {
		mylog.Error(err)
	}
	return pick.(*Unit)
}

func (this *Battle) GetAllySingleBehind(team int32) *Unit {
	front, behind := this.GetAlly(team)
	targets := behind
	if len(targets) == 0 {
		targets = front
	}
	if len(targets) == 0 {
		return nil
	}
	c := &mychooser.MyChooser{}
	for _, v := range targets {
		c.Add(v, v.GetTargetWeight())
	}
	pick, err := c.Pick()
	if err != nil {
		mylog.Error(err)
	}
	return pick.(*Unit)
}
