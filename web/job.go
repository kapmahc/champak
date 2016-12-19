package web

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// Job background job
type Job struct {
	URL       string `inject:"rabbitmq.url"`
	Namespace string `inject:"namespace"`
}

//Send send  a job message
func (p *Job) Send(queue string, body []byte) error {
	return p.open(queue, func(ch *amqp.Channel, ex, qu string) error {
		id := uuid.New().String()
		err := ch.Publish(
			ex,
			queue,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				MessageId:    id,
				Body:         body,
				Timestamp:    time.Now(),
			},
		)
		log.Infof("send message %s to %s", id, queue)
		return err
	})

}

// Receive receive a job message
func (p *Job) Receive(queue string, fn func(string, []byte, time.Time) error) error {
	return p.open(queue, func(ch *amqp.Channel, ex, qu string) error {
		msgs, err := ch.Consume(
			qu,
			"",
			true,
			false,
			false,
			false,
			nil)
		if err != nil {
			return err
		}

		for msg := range msgs {
			log.Infof("receive message %s from %s", msg.MessageId, queue)
			if err := fn(msg.MessageId, msg.Body, msg.Timestamp); err != nil {
				return err
			}
		}
		return nil
	})
}

func (p *Job) name(k string) string {
	return fmt.Sprintf("%s://%s", p.Namespace, k)
}

func (p *Job) open(queue string, fn func(*amqp.Channel, string, string) error) error {
	con, err := amqp.Dial(p.URL)
	if err != nil {
		return err
	}
	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	exn := p.name("exchange")
	if err = ch.ExchangeDeclare(exn, "direct", true, false, false, false, nil); err != nil {
		return err
	}

	qun := p.name(queue)
	if _, err := ch.QueueDeclare(qun, true, false, false, false, nil); err != nil {
		return err
	}

	if err := ch.QueueBind(qun, queue, exn, false, nil); err != nil {
		return err
	}

	return fn(ch, exn, qun)
}
