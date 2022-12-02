package myredis

import "fmt"

const (
	Account    = "Account"
	AccountUid = "AccountUid"
	CurUid     = "CurUid"
	Role       = "Role"
)

func GetRoleKey(uid uint64) string {
	return fmt.Sprintf("%s_%d", Role, uid)
}
