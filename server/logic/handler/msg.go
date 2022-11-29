package logic_handler

import (
	"gs/data/gencode"
	"gs/lib/myredis"
	"gs/proto/myproto"
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
	return &myproto.LoginACK{}
}
