package cache

import (
	"bytes"
	"encoding/gob"
	"time"
)

var (
	store Store
)

// Use use store
func Use(s Store) {
	store = s
}

// Set set
func Set(k string, v interface{}, ttl time.Duration) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return err
	}
	return store.Set(k, buf.Bytes(), ttl)
}

// Get get
func Get(k string, v interface{}) error {
	bys, err := store.Get(k)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(bys)
	return dec.Decode(v)
}

// Flush flush
func Flush() error {
	return store.Flush()
}

// Keys keys
func Keys() ([]string, error) {
	return store.Keys()
}
