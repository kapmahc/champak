package web

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Cache cache by redis
type Cache struct {
	Redis     *redis.Pool `inject:""`
	Namespace string      `inject:"namespace"`
}

//Flush clear cache items
func (p *Cache) Flush() error {
	c := p.Redis.Get()
	defer c.Close()
	keys, err := redis.Values(c.Do("KEYS", p.key("*")))
	if err == nil && len(keys) > 0 {
		_, err = c.Do("DEL", keys...)
	}
	return err
}

//Keys list cache items
func (p *Cache) Keys() ([]string, error) {
	c := p.Redis.Get()
	defer c.Close()
	return redis.Strings(c.Do("KEYS", p.key("*")))
}

// SetBytes set bytes
func (p *Cache) SetBytes(key string, val []byte, ttl time.Duration) error {
	c := p.Redis.Get()
	defer c.Close()
	_, err := c.Do("SET", p.key(key), val, "EX", int(ttl/time.Second))
	return err
}

//Set cache item
func (p *Cache) Set(key string, val interface{}, ttl time.Duration) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(val); err != nil {
		return err
	}
	return p.SetBytes(key, buf.Bytes(), ttl)
}

// GetBytes get bytes
func (p *Cache) GetBytes(key string) ([]byte, error) {
	c := p.Redis.Get()
	defer c.Close()
	return redis.Bytes(c.Do("GET", p.key(key)))
}

//Get get from cache
func (p *Cache) Get(key string, val interface{}) error {
	bys, err := p.GetBytes(key)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(bys)
	return dec.Decode(val)
}

func (p *Cache) key(k string) string {
	return fmt.Sprintf("%s@cache://%s", p.Namespace, k)
}
