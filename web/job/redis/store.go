package redis

import "time"

// Store for redis
type Store struct{}

// Send send job
func (p *Store) Send(queue string, body []byte) error {

}

// Receive recieve job
func (p *Store) Receive(queue string, fn func(string, []byte, time.Time) error) error {

}
