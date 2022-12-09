package myredis

import "fmt"

const (
	Account      = "Account"
	AccountUid   = "AccountUid"
	CurUid       = "CurUid"
	Role         = "Role"
	RoleName     = "RoleName"
	NotifyPlayer = "notify_player"
)

func GetRoleKey(uid uint64) string {
	return fmt.Sprintf("%s_%d", Role, uid)
}
