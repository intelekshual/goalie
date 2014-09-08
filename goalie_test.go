package goalie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMemoryProvider(t *testing.T) {
	assert := assert.New(t)

	acl := NewMemoryProvider()

	ok, _ := acl.Assert("me", "job-1234")
	assert.Equal(ok, false, "Should be false")

	acl.Grant("me", "job-1234")
	ok, _ = acl.Assert("me", "job-1234")
	assert.Equal(ok, true, "Should be true")

	acl.Revoke("me", "job-1234")
	ok, _ = acl.Assert("me", "job-1234")
	assert.Equal(ok, false, "Should be false")
}

func TestNewRedisProvider(t *testing.T) {
	assert := assert.New(t)

	acl := NewRedisProvider(map[string]string{})

	ok, _ := acl.Assert("me", "job-1234")
	assert.Equal(ok, false, "Should be false")

	acl.Grant("me", "job-1234")
	ok, _ = acl.Assert("me", "job-1234")
	assert.Equal(ok, true, "Should be true")

	acl.Revoke("me", "job-1234")
	ok, _ = acl.Assert("me", "job-1234")
	assert.Equal(ok, false, "Should be false")
}
