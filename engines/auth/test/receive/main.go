package main

import (
	"fmt"
	"log"

	"github.com/kapmahc/champak/web"
)

func sendMail(t, s, b string) (string, error) {
	fmt.Printf("send mail to %s:\n%s\n%s\n", t, s, b)
	return "done", nil
}

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

func main() {
	srv, err := web.NewWorkerServer()
	if err != nil {
		log.Fatal(err)
	}

	// -------
	srv.RegisterTasks(map[string]interface{}{
		"add":             add,
		"multiply":        multiply,
		"auth.send-email": sendMail,
	})

	if err := srv.NewWorker("test").Launch(); err != nil {
		log.Fatal(err)
	}
}
