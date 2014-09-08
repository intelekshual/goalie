package goalie

type PermissionsProvider interface {
	Grant(g string, r string) error
	Assert(g string, r string) (bool, error)
	Revoke(g string, r string) error
}
