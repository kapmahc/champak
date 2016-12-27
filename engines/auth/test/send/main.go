package main

import (
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v1/signatures"
	"github.com/kapmahc/champak/web"
)

func main() {
	srv, err := web.NewWorkerServer()
	if err != nil {
		log.Fatal(err)
	}

	// --------------------

	for i := 0; i < 10; i++ {
		task := signatures.TaskSignature{
			Name: "auth.send-email",
			Args: []signatures.TaskArg{
				signatures.TaskArg{
					Type:  "string",
					Value: fmt.Sprintf("a%d@aaa.com", i),
				},
				signatures.TaskArg{
					Type:  "string",
					Value: "sss",
				},
				signatures.TaskArg{
					Type:  "string",
					Value: "bbb",
				},
			},
		}

		if _, err := srv.SendTask(&task); err != nil {
			log.Fatal(err)
		}
	}

	// ----------------------

}
