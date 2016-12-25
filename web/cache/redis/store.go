package redis

import (
	"fmt"
	"time"

	_redis "github.com/garyburd/redigo/redis"
)

// Store cache store by redis
type Store struct {
	pool   *_redis.Pool
	prefix string
}

// Use use redis pool
func Use(pool *_redis.Pool, prefix string) *Store {
	return &Store{
		pool:   pool,
		prefix: prefix,
	}
}

// New open redis connection
func New(host string, port, db int, prefix string) *Store {
	return &Store{
		prefix: prefix,
		pool: &_redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (_redis.Conn, error) {
				c, e := _redis.Dial(
					"tcp",
					fmt.Sprintf(
						"%s:%d",
						host,
						port,
					),
				)
				if e != nil {
					return nil, e
				}
				if _, e = c.Do("SELECT", db); e != nil {
					c.Close()
					return nil, e
				}
				return c, nil
			},
			TestOnBorrow: func(c _redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}

}

//Flush flush
func (p *Store) Flush() error {
	c := p.pool.Get()
	defer c.Close()
	keys, err := _redis.Values(c.Do("KEYS", p.key("*")))
	if err == nil && len(keys) > 0 {
		_, err = c.Do("DEL", keys...)
	}
	return err
}

//Keys list cache items
func (p *Store) Keys() ([]string, error) {
	c := p.pool.Get()
	defer c.Close()
	return _redis.Strings(c.Do("KEYS", p.key("*")))
}

//Set set
func (p *Store) Set(k string, v []byte, ttl time.Duration) error {
	c := p.pool.Get()
	defer c.Close()
	_, err := c.Do("SET", p.key(k), v, "EX", int(ttl/time.Second))
	return err
}

//Get get
func (p *Store) Get(k string) ([]byte, error) {
	c := p.pool.Get()
	defer c.Close()
	return _redis.Bytes(c.Do("GET", p.key(k)))
}

func (p *Store) key(k string) string {
	return fmt.Sprintf("cache://%s/%s", p.prefix, k)
}
