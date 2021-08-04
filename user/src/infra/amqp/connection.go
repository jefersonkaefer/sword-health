package amqp

import (
	"fmt"
	"log"
	"sword-health/users/application/command"

	"github.com/streadway/amqp"
)

type Connection struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	msgs    <-chan amqp.Delivery
}

func (Connection) New(user string, password string, host string, port int) *Connection {
	var err error

	r := Connection{}

	addr := fmt.Sprintf(
		"amqp://%s:%s@%s:%v/",
		user,
		password,
		host,
		port,
	)

	r.conn, err = amqp.Dial(addr)

	if err != nil {
		log.Println("ERROR::: ", err)
	}

	r.channel, err = r.conn.Channel()

	if err != nil {
		log.Println("ERROR::: ", err)
	}

	return &r
}

func (r *Connection) DeclareExchange(exchangeName string) *Connection {
	var err error

	err = r.channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("ERROR::: ", err.Error())
	}

	return r
}

func (r *Connection) QueueDeclare(exchange string, queueName string, routeKey string) *Connection {

	queue, err := r.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		true,
		nil,
	)

	if err != nil {
		log.Println("ERROR::: ", err.Error())
	}

	err = r.channel.QueueBind(
		queue.Name, // queue name
		routeKey,   // routing key
		exchange,   // exchange
		false,
		nil,
	)

	return r
}
func (r *Connection) Consume(consumerName string, cmd *command.UserHandler, queueName string) {

	if r.conn == nil {
		panic("Connection error.")
	}

	ch, err := r.conn.Channel()

	if err != nil {
		log.Println("ERROR::: ", err.Error())
	}
	defer ch.Close()

	r.msgs, err = ch.Consume(
		queueName,    // queue
		consumerName, // consumer
		true,         // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)

	if err != nil {
		log.Println("ERROR::: ", err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range r.msgs {
			cmd.Exec(d.RoutingKey, d.Body)
		}
	}()

	log.Printf(" [*] AMQP Ready")
	<-forever
}
