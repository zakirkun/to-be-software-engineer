package rabbitmq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Address string
}

// Publish implements IRabbitMQ.
func (r RabbitMQ) Publish(queueName string, body interface{}) error {
	conn, err := amqp.Dial(r.Address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		return err
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key (queue name)
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json", // Set ContentType to JSON
			Body:        jsonBody,           // Send JSON body
		})

	if err != nil {
		return err
	}

	log.Printf(" [x] Sent JSON message: %s", jsonBody)

	return nil
}

// Listener implements IRabbitMQ.
func (r RabbitMQ) Listener(queueName string, cb ...func(payload interface{}) error) {
	conn, err := amqp.Dial(r.Address)
	failOnError(err, "Failed open connection")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare the same queue
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			for _, f := range cb {
				f(d.Body)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// Open implements IRabbitMQ.
func (r RabbitMQ) Open() (*amqp.Connection, error) {
	conn, err := amqp.Dial(r.Address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn, nil
}

type IRabbitMQ interface {
	Open() (*amqp.Connection, error)
	Listener(
		queueName string,
		cb ...func(payload interface{}) error,
	)
	Publish(queueName string, body interface{}) error
}

func NewRabbitMQ(addr string) IRabbitMQ {
	return RabbitMQ{
		Address: addr,
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
