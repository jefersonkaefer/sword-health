package amqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Connection struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewConnection(user string, password string, host string, port int) *Connection {
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
		fmt.Println("ERROR::: ", err)
	}

	r.channel, err = r.conn.Channel()

	if err != nil {
		fmt.Println("ERROR::: ", err)
	}
	fmt.Println("me contactando...", r)
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
		fmt.Println("ERROR::: ", err.Error())
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
		fmt.Println("ERROR::: ", err.Error())
	}

	err = r.channel.QueueBind(
		queue.Name, // queue name
		routeKey,   // routing key
		exchange,   // exchange
		false,
		nil,
	)

	if err != nil {
		fmt.Println("ERROR::: ", err.Error())
	}

	return r
}
func (r *Connection) Dispatch(exchangeName string, routeKey string, message []byte) error {
	if r.conn == nil {
		panic("Connection error.")
	}

	ch, err := r.conn.Channel()

	if err != nil {
		fmt.Println("ERROR::: ", err.Error())
	}
	defer ch.Close()

	return ch.Publish(
		exchangeName,
		routeKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/json",
			Body:         message,
			DeliveryMode: amqp.Persistent,
		})
}

func (r *Connection) Close() {
	r.channel.Close()
	r.conn.Close()
}
