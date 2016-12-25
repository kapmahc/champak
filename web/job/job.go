package job

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

var (
	store Store
)

// Use use store
func Use(s Store) {
	store = s
}

// Send send a job message
func Send(queue string, body []byte) {
	if err := store.Send(queue, body); err != nil {
		log.Error(err)
	}
}

// Receive recieve a job message and do it
func Receive(queue string, fn func(string, []byte, time.Time) error) {
	if err := store.Receive(queue, fn); err != nil {
		log.Error(err)
	}
}
