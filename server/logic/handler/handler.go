package logic_handler

import (
	"gs/lib/mylog"
	"gs/proto/myproto"
)

func ProcessMsg(uid uint64, msgId uint32, data []byte) (myproto.MsgId, myproto.MyMsg) {
	switch myproto.MsgId(msgId) {
	case myproto.MsgId_Msg_RegisterREQ:
		return myproto.MsgId_Msg_RegisterACK, handleRegister(uid, data)
	case myproto.MsgId_Msg_LoginREQ:
		return myproto.MsgId_Msg_LoginACK, handleLogin(uid, data)
	default:
		mylog.Error("msg id ", msgId, " not handle")
	}
	return 0, nil
}