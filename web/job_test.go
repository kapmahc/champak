package web_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/kapmahc/champak/web"
)

func TestJob(t *testing.T) {
	job := web.Job{URL: "amqp://guest:guest@localhost:5672/", Namespace: "champak-test"}
	job.Send("echo", []byte("Hello, Job!"))

	if err := job.Receive("echo", func(id string, body []byte, created time.Time) error {
		fmt.Printf("ECHO [%s] %s %v\n", id, body, created)
		return errors.New("normal return")
	}); err != nil {
		// t.Fatal(err)
		t.Log(err)
	}
}
