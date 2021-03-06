// Package queue - for amqp queue
package queue

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

func errorHandle(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Exchange is struct
type Exchange struct {
	Name       string
	Channel    *amqp.Channel
	DelayQueue map[string]amqp.Queue
}

// BindQueue is bind queue to exchange
func (e *Exchange) BindQueue(name, key string) {
	err := e.Channel.QueueBind(
		name,   //name
		key,    //key
		e.Name, //exchange
		false,  //noWait
		nil,    //args
	)
	errorHandle(err, fmt.Sprintf("QueueBind `%s` fail\n", name))
}

// DeclareDeadLetter is to create delay queue
func (e *Exchange) DeclareDeadLetter(name, routingKey string) {
	args := amqp.Table{
		"x-dead-letter-exchange":    e.Name,
		"x-dead-letter-routing-key": routingKey,
	}
	queueDelay, err := e.Channel.QueueDeclare(name, true, false, false, false, args)
	errorHandle(err, fmt.Sprintf("Can't declare `%s` queue\n", name))
	e.DelayQueue[name] = queueDelay
}

// Publish is to publish
func (e *Exchange) Publish(key string, body []byte) {
	err := e.Channel.Publish(
		e.Name, //exchange
		key,    //key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Timestamp:    time.Now(),
			Body:         body,
			MessageId:    uuid.NewV4().String(),
		})
	errorHandle(err, "Publish fail")
	if err != nil {
		log.Fatalln(err)
	}
}

// MQ is a struct
type MQ struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	Exchange map[string]*Exchange
}

// Close connection
func (m *MQ) Close() {
	m.Channel.Close()
	m.Conn.Close()
}

// SetQos to set channel qos
func (m *MQ) SetQos(prefetchCount int) {
	m.Channel.Qos(prefetchCount, 0, true)
}

// DeclareExchange an Exchange
func (m *MQ) DeclareExchange(name string) {
	err := m.Channel.ExchangeDeclare(
		name,               //name
		amqp.ExchangeTopic, //kind
		true,               //durable
		false,              //autoDelete
		false,              //internal
		false,              //noWait
		nil,                //args
	)
	errorHandle(err, fmt.Sprintf("Can't declare `%s` exchange\n", name))

	m.Exchange[name] = &Exchange{
		Name:       name,
		Channel:    m.Channel,
		DelayQueue: make(map[string]amqp.Queue),
	}
}

// DeclareQueue an queue
func (m *MQ) DeclareQueue(name string) {
	queue, err := m.Channel.QueueDeclare(
		name,  //name
		true,  //durable
		false, //autoDelete
		false, //exclusive
		false, //noWait
		nil,   //args
	)
	errorHandle(err, fmt.Sprintf("Can't declare `%s` queue", name))
	log.Printf("DeclareQueue: %+v\n", queue)
}

// GetConsumer is to create consumer
func (m *MQ) GetConsumer(name string) <-chan amqp.Delivery {
	queueConsumer, err := m.Channel.Consume(
		name,  //queue
		"",    //consumer
		false, //autoAck
		false, //exclusive
		false, //noLocal
		false, //noWait
		nil,   //args
	)
	errorHandle(err, fmt.Sprintf("Can't register consumer `%s`", name))
	return queueConsumer
}

// New a MQ struct
func New(url string) *MQ {
	conn, err := amqp.Dial(url)
	errorHandle(err, "RabbitMQ conn fail")

	channel, err := conn.Channel()
	errorHandle(err, "Establish channel fail")

	channel.Qos(32, 0, false)

	return &MQ{
		Conn:     conn,
		Channel:  channel,
		Exchange: make(map[string]*Exchange),
	}
}
