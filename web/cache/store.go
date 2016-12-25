package cache

import "time"

// Store cache store
type Store interface {
	Flush() error
	Keys() ([]string, error)
	Set(k string, v []byte, ttl time.Duration) error
	Get(k string) ([]byte, error)
}
