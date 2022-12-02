package logic_handler

import (
	"encoding/json"
	"gs/data/gencode"
	"gs/lib/myredis"
	"gs/lib/myutil"
	"gs/proto/myproto"
	"gs/server/logic/player"
	"strconv"
	"time"
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
	player := &player.Player{
		Uid:      uid,
		Name:     req.Name,
		CreateAt: time.Now().Unix(),
	}
	jsonData, err := json.Marshal(player)
	if err != nil {
		return &myproto.CreateRoleACK{Ret: myproto.ResultCode_CreateRoleFaild}
	}
	if ok = myredis.GetInstance().Set(myredis.GetRoleKey(uid), jsonData, 0); !ok {
		return &myproto.CreateRoleACK{Ret: myproto.ResultCode_CreateRoleFaild}
	}
	return &myproto.CreateRoleACK{}
}
