package web

import (
	"fmt"

	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/logger"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewWorkerServer new machinery server
func NewWorkerServer() (*machinery.Server, error) {
	url := fmt.Sprintf(
		"redis://%s:%d/%d",
		viper.GetString("redis.host"),
		viper.GetInt("redis.port"),
		viper.GetInt("redis.db"),
	)
	return machinery.NewServer(&config.Config{
		Broker:          url,
		ResultBackend:   url,
		ResultsExpireIn: 60 * 60 * 24 * 30 * 6,
		DefaultQueue:    fmt.Sprintf("tasks://%s", viper.GetString("app.name")),
		// Broker:        "amqp://guest:guest@localhost:5672/",
		// ResultBackend: "amqp://guest:guest@localhost:5672/",
		// Exchange:      "machinery_exchange",
		// ExchangeType:  "direct",
		// DefaultQueue:  "machinery_tasks",
		// BindingKey:    "machinery_task",
	})
}

type machineryLogger struct {
}

func (p *machineryLogger) Print(args ...interface{}) {
	log.Print(args...)
}
func (p *machineryLogger) Printf(f string, args ...interface{}) {
	log.Printf(f, args...)
}
func (p *machineryLogger) Println(args ...interface{}) {
	log.Println(args...)
}

func (p *machineryLogger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}
func (p *machineryLogger) Fatalf(f string, args ...interface{}) {
	log.Fatalf(f, args...)
}
func (p *machineryLogger) Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}

func (p *machineryLogger) Panic(args ...interface{}) {
	log.Panic(args...)
}
func (p *machineryLogger) Panicf(f string, args ...interface{}) {
	log.Panicf(f, args...)
}
func (p *machineryLogger) Panicln(args ...interface{}) {
	log.Panicln(args...)
}

func init() {
	logger.Set(&machineryLogger{})
}
