package redis

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	_redis "github.com/garyburd/redigo/redis"
	"github.com/google/uuid"
)

// Store for redis
type Store struct {
	pool   *_redis.Pool
	prefix string
}

// Message job message
type Message struct {
	ID      string
	Body    []byte
	Created time.Time
}

// Send send job
func (p *Store) Send(queue string, body []byte) error {
	msg := Message{
		ID:      uuid.New().String(),
		Body:    body,
		Created: time.Now(),
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(msg); err != nil {
		return err
	}

	c := p.pool.Get()
	defer c.Close()
	_, e := c.Do("LPUSH", p.name(queue), buf.Bytes())
	return e
}

// Receive recieve job
func (p *Store) Receive(queue string, fn func(string, []byte, time.Time) error) error {
	c := p.pool.Get()
	defer c.Close()
	for {
		body, err := _redis.Bytes(c.Do("RPOP", p.name(queue)))
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		dec := gob.NewDecoder(&buf)
		buf.Write(body)
		var msg Message
		if err := dec.Decode(&msg); err != nil {
			return err
		}

		if err := fn(msg.ID, msg.Body, msg.Created); err != nil {
			return err
		}
	}

	return nil
}

func (p *Store) name(q string) string {
	return fmt.Sprintf("task://%s/%s", p.prefix, q)
}
