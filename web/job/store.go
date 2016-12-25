package job

import "time"

// Store store
type Store interface {
	Send(queue string, body []byte) error
	Receive(queue string, fn func(string, []byte, time.Time) error) error
}
