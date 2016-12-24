package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var namespace string
var pool *redis.Pool

// Open open redis connection
func Open(host string, port, db int, ns string) {
	namespace = ns
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, e := redis.Dial(
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
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

//Flush clear cache items
func Flush() error {
	c := pool.Get()
	defer c.Close()
	keys, err := redis.Values(c.Do("KEYS", key("*")))
	if err == nil && len(keys) > 0 {
		_, err = c.Do("DEL", keys...)
	}
	return err
}

//Keys list cache items
func Keys() ([]string, error) {
	c := pool.Get()
	defer c.Close()
	return redis.Strings(c.Do("KEYS", key("*")))
}

//Set cache item
func Set(k string, v interface{}, ttl uint) error {
	c := pool.Get()
	defer c.Close()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return err
	}
	_, err := c.Do("SET", key(k), buf.Bytes(), "EX", ttl)
	return err
}

//Get get from cache
func Get(k string, v interface{}) error {
	c := pool.Get()
	defer c.Close()
	bys, err := redis.Bytes(c.Do("GET", key(k)))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(bys)
	return dec.Decode(v)
}

func key(k string) string {
	return fmt.Sprintf("cache://%s/%s", namespace, k)
}
