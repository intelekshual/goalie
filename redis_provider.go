package goalie

import (
	"github.com/garyburd/redigo/redis"
)

type RedisProvider struct {
	prefix  string
	network string
	address string
	conn    redis.Conn
}

type redisConfig map[string]string

func NewRedisProvider(config redisConfig) *RedisProvider {
	rp := &RedisProvider{
		prefix:  "goalie:",
		network: "tcp",
		address: ":6379",
	}

	// overwrite defaults if values are present in the
	// provided config map
	for key, value := range config {
		switch key {
		case "prefix":
			rp.prefix = value
		case "network":
			rp.network = value
		case "address":
			rp.address = value
		}
	}

	conn, err := redis.Dial(rp.network, rp.address)
	if err != nil {
		panic("Could not connect to Redis")
	}
	rp.conn = conn

	// ensures interface is checked
	var _ PermissionsProvider = rp
	return rp
}

func (rp *RedisProvider) Grant(g, r string) error {
	_, err := rp.conn.Do("SADD", rp.key(g), r)
	return err
}

func (rp *RedisProvider) Assert(g, r string) (bool, error) {
	ok, err := redis.Bool(rp.conn.Do("SISMEMBER", rp.key(g), r))
	return ok, err
}

func (rp *RedisProvider) Revoke(g, r string) error {
	_, err := rp.conn.Do("SREM", rp.key(g), r)
	return err
}

// Returns a prefixed key
func (rp *RedisProvider) key(g string) string {
	return rp.prefix + g
}
