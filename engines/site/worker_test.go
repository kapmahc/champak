package site_test

import (
	"fmt"
	"testing"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/signatures"
)

func add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	fmt.Println(sum)
	return sum, nil
}

func multiply(args ...int64) (int64, error) {
	sum := int64(1)
	for _, arg := range args {
		sum *= arg
	}
	fmt.Println(sum)
	return sum, nil
}

func TestWorker(t *testing.T) {
	url := "redis://localhost:6379"
	srv, err := machinery.NewServer(&config.Config{
		Broker:          url,
		ResultBackend:   url,
		ResultsExpireIn: 60 * 60 * 24 * 30 * 100,
		DefaultQueue:    "test://tasks/",
	})

	if err != nil {
		t.Fatal(err)
	}
	// -------
	srv.RegisterTasks(map[string]interface{}{
		"add":      add,
		"multiply": multiply,
	})

	// --------------------

	for i := 0; i < 10; i++ {
		task := signatures.TaskSignature{
			Name: "add",
			Args: []signatures.TaskArg{
				signatures.TaskArg{
					Type:  "int64",
					Value: i,
				},
				signatures.TaskArg{
					Type:  "int64",
					Value: 100,
				},
			},
		}

		if _, err := srv.SendTask(&task); err != nil {
			t.Fatal(err)
		}
	}

	// ----------------------

	//  if err := srv.NewWorker("test").Launch(); err != nil {
	//  	t.Fatal(err)
	//  }
}
