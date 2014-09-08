package goalie

import (
	"github.com/deckarep/golang-set"
)

type MemoryProvider struct {
	permissions map[string]mapset.Set
}

func NewMemoryProvider() *MemoryProvider {
	mp := &MemoryProvider{}
	mp.permissions = make(map[string]mapset.Set)
	var _ PermissionsProvider = mp
	return mp
}

func (mp *MemoryProvider) Grant(g, r string) error {
	if mp.permissions[g] == nil {
		mp.permissions[g] = mapset.NewSet()
	}
	mp.permissions[g].Add(r)
	return nil
}

func (mp *MemoryProvider) Assert(g, r string) (bool, error) {
	if mp.permissions[g] == nil {
		return false, nil
	}

	return mp.permissions[g].Contains(r), nil
}

func (mp *MemoryProvider) Revoke(g, r string) error {
	if mp.permissions[g] != nil {
		mp.permissions[g].Remove(r)
	}
	return nil
}
