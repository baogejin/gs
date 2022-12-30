package logic_handler

import (
	"encoding/json"
	"gs/data/gencode"
	"gs/define"
	"gs/game/battle"
	"gs/game/player_info"
	"gs/lib/mylog"
	"gs/lib/myredis"
	"gs/lib/myrpc"
	"gs/lib/myutil"
	"gs/proto/myproto"
	"gs/server/logic/battle_manager"
	"gs/server/logic/player_manager"
	"strconv"
	"strings"
)

func handleRegister(uid uint64, data []byte) *myproto.RegisterACK {
	if uid > 0 {
		return &myproto.RegisterACK{Ret: myproto.ResultCode_AlreadyLogin}
	}
	req := &myproto.RegisterREQ{}
	err := req.Unmarshal(data)
	if err != nil {
		return &myproto.RegisterACK{Ret: myproto.ResultCode_MsgErr}
	}
	if req.Account == "" {
		return &myproto.RegisterACK{Ret: myproto.ResultCode_AccountEmpty}
	}
	if req.Password == "" {
		return &myproto.RegisterACK{Ret: myproto.ResultCode_PasswordEmpty}
	}
	if info, ok := gencode.GetGlobalCfg().GetGlobalInfoByKey(gencode.AccountMaxLen); ok {
		if len(req.Account) > int(info.Value) {
			return &myproto.RegisterACK{Ret: myproto.ResultCode_AccountErr}
		}
	}
	if len(req.Account) != len([]rune(req.Account)) {
		return &myproto.RegisterACK{Ret: myproto.ResultCode_AccountErr}
	}
	ok := myredis.GetInstance().HSetNX(myredis.Account, req.Account, req.Password)
	if !ok {
		return &myproto.RegisterACK{Ret: myproto.ResultCode_AccountExist}
	}
	return &myproto.RegisterACK{}
}

func handleLogin(uid uint64, data []byte) *myproto.LoginACK {
	if uid > 0 {
		return &myproto.LoginACK{Ret: myproto.ResultCode_AlreadyLogin}
	}
	req := &myproto.RegisterREQ{}
	err := req.Unmarshal(data)
	if err != nil {
		return &myproto.LoginACK{Ret: myproto.ResultCode_MsgErr}
	}
	pwd := myredis.GetInstance().HGet(myredis.Account, req.Account)
	if pwd == "" {
		return &myproto.LoginACK{Ret: myproto.ResultCode_AccountNotExist}
	}
	if pwd != req.Password {
		return &myproto.LoginACK{Ret: myproto.ResultCode_PasswordErr}
	}
	uidStr := myredis.GetInstance().HGet(myredis.AccountUid, req.Account)
	uid, err = strconv.ParseUint(uidStr, 10, 64)
	if err != nil || uid == 0 {
		uid = uint64(myredis.GetInstance().Incr(myredis.CurUid))
		myredis.GetInstance().HSet(myredis.AccountUid, req.Account, uid)
	}
	ok := myredis.GetInstance().Exist(myredis.GetRoleKey(uid))
	return &myproto.LoginACK{Uid: uid, HasRole: ok}
}

func handCreateRole(uid uint64, data []byte) *myproto.CreateRoleACK {
	if uid == 0 {
		return &myproto.CreateRoleACK{Ret: myproto.ResultCode_NeedLogin}
	}
	req := &myproto.CreateRoleREQ{}
	err := req.Unmarshal(data)
	if err != nil {
		return &myproto.CreateRoleACK{Ret: myproto.ResultCode_MsgErr}
	}
	if myredis.GetInstance().Exist(myredis.GetRoleKey(uid)) {
		return &myproto.CreateRoleACK{Ret: myproto.ResultCode_AlreadyHasRole}
	}
	nameLen := myutil.GetStringLen(req.Name)
	if cfg, ok := gencode.GetGlobalCfg().GetGlobalInfoByKey(gencode.RoleNameMinLen); ok {
		if nameLen < int(cfg.Value) {
			return &myproto.CreateRoleACK{Ret: myproto.ResultCode_RoleNameIllegal}
		}
	}
	if cfg, ok := gencode.GetGlobalCfg().GetGlobalInfoByKey(gencode.RoleNameMaxLen); ok {
		if nameLen > int(cfg.Value) {
			return &myproto.CreateRoleACK{Ret: myproto.ResultCode_RoleNameIllegal}
		}
	}
	ok := myredis.GetInstance().HSet(myredis.RoleName, req.Name, uid)
	if !ok {
		return &myproto.CreateRoleACK{Ret: myproto.ResultCode_RoleNameExist}
	}
	player := player_info.NewPlayer(uid, req.Name)
	ok = player.Save()
	if !ok {
		return &myproto.CreateRoleACK{Ret: myproto.ResultCode_CreateRoleFaild}
	}
	return &myproto.CreateRoleACK{}
}

func handEnterGame(uid uint64, notifyAddr string) *myproto.EnterGameACK {
	player := player_manager.GetMgr().GetPlayer(uid)
	if player == nil {
		player = player_info.NewPlayer(0, "")
		jsonData := myredis.GetInstance().Get(myredis.GetRoleKey(uid))
		if len(jsonData) == 0 {
			return &myproto.EnterGameACK{Ret: myproto.ResultCode_EnterGameFailed}
		}
		err := json.Unmarshal([]byte(jsonData), player)
		if err != nil {
			mylog.Error(err)
			return &myproto.EnterGameACK{Ret: myproto.ResultCode_EnterGameFailed}
		}
		player_manager.GetMgr().SetPlayer(uid, player)
	}
	player.SetNotifyAddr(notifyAddr)
	// player.SendMsg(myproto.MsgId_Msg_EnterGameACK, &myproto.EnterGameACK{}) //notify test
	// player_manager.GetMgr().BroadcastAllPlayer(myproto.MsgId_Msg_EnterGameACK, &myproto.EnterGameACK{Info: &myproto.PlayerInfo{Uid: 11111, Name: "2222"}})
	return &myproto.EnterGameACK{Info: player.Proto()}
}

func handleLogout(uid uint64) *myproto.LogoutACK {
	player := player_manager.GetMgr().GetPlayer(uid)
	if player != nil {
		//todo 其他下线处理
		if player.Save() {
			player_manager.GetMgr().DelPlayer(uid)
		}
	}
	return &myproto.LogoutACK{}
}

func handChat(uid uint64, data []byte) *myproto.ChatACK {
	player := player_manager.GetMgr().GetPlayer(uid)
	if player == nil {
		return &myproto.ChatACK{Ret: myproto.ResultCode_PlayerNotFound}
	}
	req := &myproto.ChatREQ{}
	err := req.Unmarshal(data)
	if err != nil {
		return &myproto.ChatACK{Ret: myproto.ResultCode_MsgErr}
	}
	push := &myproto.ChatPUSH{
		Uid:  player.Uid,
		Name: player.Name,
		Msg:  req.Msg,
	}
	myrpc.GetInstance().NotifyAllNodes(define.NodeGateway, myproto.MsgId_Msg_ChatPUSH, push)
	return &myproto.ChatACK{}
}

func handGM(uid uint64, data []byte) *myproto.GMACK {
	player := player_manager.GetMgr().GetPlayer(uid)
	if player == nil {
		return &myproto.GMACK{Ret: myproto.ResultCode_PlayerNotFound}
	}
	req := &myproto.GMREQ{}
	err := req.Unmarshal(data)
	if err != nil {
		return &myproto.GMACK{Ret: myproto.ResultCode_MsgErr}
	}
	args := strings.Split(req.Cmd, " ")
	if len(args) == 0 {
		return &myproto.GMACK{Ret: myproto.ResultCode_GMCmdNotFound}
	}
	switch strings.ToLower(args[0]) {
	case "additem":
		if len(args) != 3 {
			return &myproto.GMACK{Ret: myproto.ResultCode_GMCmdParamErr}
		}
		itemid, err := strconv.ParseInt(args[1], 10, 0)
		if err != nil {
			return &myproto.GMACK{Ret: myproto.ResultCode_GMCmdParamErr}
		}
		num, err := strconv.ParseInt(args[2], 10, 0)
		if err != nil {
			return &myproto.GMACK{Ret: myproto.ResultCode_GMCmdParamErr}
		}
		player.AddItems(&myproto.Item{ItemId: int32(itemid), Num: num})
		return &myproto.GMACK{}
	default:
		return &myproto.GMACK{Ret: myproto.ResultCode_GMCmdNotFound}
	}
}

func handCreateBattle(uid uint64, data []byte) *myproto.CreateBattleACK {
	player := player_manager.GetMgr().GetPlayer(uid)
	if player == nil {
		return &myproto.CreateBattleACK{Ret: myproto.ResultCode_PlayerNotFound}
	}
	req := &myproto.CreateBattleREQ{}
	err := req.Unmarshal(data)
	if err != nil {
		return &myproto.CreateBattleACK{Ret: myproto.ResultCode_MsgErr}
	}
	b := battle.CreatePveBattle(req.LevelId, player.GenBattleUnit())
	if b == nil {
		return &myproto.CreateBattleACK{Ret: myproto.ResultCode_CreateBattleFailed}
	}
	b.BattleInfoNotify()
	battle_manager.GetMgr().AddBattle(b)
	// b.Start()
	return &myproto.CreateBattleACK{}
}

func handBattleStart(uid uint64, data []byte) *myproto.BattleStartACK {
	player := player_manager.GetMgr().GetPlayer(uid)
	if player == nil {
		return &myproto.BattleStartACK{Ret: myproto.ResultCode_PlayerNotFound}
	}
	req := &myproto.BattleStartREQ{}
	err := req.Unmarshal(data)
	if err != nil {
		return &myproto.BattleStartACK{Ret: myproto.ResultCode_MsgErr}
	}
	b := battle_manager.GetMgr().GetBattle(req.BattleId)
	if b != nil {
		b.Start()
	}
	return &myproto.BattleStartACK{}
}

func handleBattleSkill(uid uint64, data []byte) *myproto.BattleSkillACK {
	player := player_manager.GetMgr().GetPlayer(uid)
	if player == nil {
		return &myproto.BattleSkillACK{Ret: myproto.ResultCode_PlayerNotFound}
	}
	req := &myproto.BattleSkillREQ{}
	err := req.Unmarshal(data)
	if err != nil {
		return &myproto.BattleSkillACK{Ret: myproto.ResultCode_MsgErr}
	}
	b := battle_manager.GetMgr().GetBattle(req.BattleId)
	if b != nil {
		b.SetNextSkill(uid, req.SkillId)
	}
	return &myproto.BattleSkillACK{}
}
