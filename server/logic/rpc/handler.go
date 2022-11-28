package rpc_logic

import (
	"gs/lib/mylog"
	"gs/lib/myredis"
	"gs/proto/myproto"
)

func processMsg(uid uint64, msgId uint32, data []byte) (myproto.MsgId, myproto.MyMsg) {
	switch myproto.MsgId(msgId) {
	case myproto.MsgId_Msg_RegisterREQ:
		return myproto.MsgId_Msg_RegisterACK, handleRegister(uid, data)
	default:
		mylog.Error("msg id ", msgId, " not handle")
	}
	return 0, nil
}

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
	ok := myredis.GetInstance().HSetNX(myredis.Account, req.Account, req.Password)
	if !ok {
		return &myproto.RegisterACK{Ret: myproto.ResultCode_AccountExist}
	}
	return &myproto.RegisterACK{}
}
