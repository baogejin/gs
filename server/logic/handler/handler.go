package logic_handler

import (
	"gs/lib/mylog"
	"gs/proto/myproto"
)

func ProcessMsg(uid uint64, msgId uint32, data []byte, notifyAddr string) (myproto.MsgId, myproto.MyMsg) {
	switch myproto.MsgId(msgId) {
	case myproto.MsgId_Msg_RegisterREQ:
		return myproto.MsgId_Msg_RegisterACK, handleRegister(uid, data)
	case myproto.MsgId_Msg_LoginREQ:
		return myproto.MsgId_Msg_LoginACK, handleLogin(uid, data)
	case myproto.MsgId_Msg_CreateRoleREQ:
		return myproto.MsgId_Msg_CreateRoleACK, handCreateRole(uid, data)
	case myproto.MsgId_Msg_EnterGameREQ:
		return myproto.MsgId_Msg_EnterGameACK, handEnterGame(uid, notifyAddr)
	case myproto.MsgId_Msg_LogoutREQ:
		return myproto.MsgId_Msg_LogoutACK, handleLogout(uid)
	default:
		mylog.Error("msg id ", msgId, " not handle")
	}
	return 0, nil
}
