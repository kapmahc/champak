package amqp

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	_amqp "github.com/streadway/amqp"
)

// New new amqp store
func New(host string, port int, user, password, ns string) *Store {
	return &Store{
		url: fmt.Sprintf("amqp://%s:%s@%s:%d/", user, password, host, port),
		ns:  ns,
	}
}

// Store for store
type Store struct {
	url string
	ns  string
}

// Send send job
func (p *Store) Send(queue string, body []byte) error {
	return p.open(queue, func(ch *_amqp.Channel, ex, qu string) error {
		id := uuid.New().String()
		log.Infof("send message %s to %s", id, queue)
		return ch.Publish(
			ex,
			queue,
			false,
			false,
			_amqp.Publishing{
				DeliveryMode: _amqp.Persistent,
				ContentType:  "text/plain",
				MessageId:    id,
				Body:         body,
				Timestamp:    time.Now(),
			},
		)
	})
}

// Receive recieve job
func (p *Store) Receive(queue string, fn func(string, []byte, time.Time) error) error {
	return p.open(queue, func(ch *_amqp.Channel, ex, qu string) error {
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

func (p *Store) name(k string) string {
	return fmt.Sprintf("%s://%s", p.ns, k)
}

func (p *Store) open(queue string, fn func(*_amqp.Channel, string, string) error) error {
	con, err := _amqp.Dial(p.url)
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
