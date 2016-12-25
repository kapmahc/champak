package job_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/kapmahc/champak/web/job"
	"github.com/kapmahc/champak/web/job/amqp"
)

func TestRabbitMQ(t *testing.T) {
	s := amqp.New("localhost", 5672, "guest", "guest", "test")
	testJob(t, s)
}

func testJob(t *testing.T, s job.Store) {
	job.Use(s)
	job.Send("echo", []byte("Hello, Job!"))

	job.Receive("echo", func(id string, body []byte, created time.Time) error {
		fmt.Printf("ECHO [%s] %s %v\n", id, body, created)
		return errors.New("normal return")
	})
}
