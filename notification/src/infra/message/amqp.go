package message

import (
	"log"
	"sword-health/notification/application/command"

	"github.com/streadway/amqp"
)

type AMQP struct {
	dsn     string
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (AMQP) New(dsn string) Broker {

	return &AMQP{
		dsn: dsn,
	}
}

func (b *AMQP) QueueDeclare(exchange string, queueName string, routeKey string) Broker {
	var err error
	b.connect()
	err = b.channel.ExchangeDeclare(
		exchange,
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

	queue, err := b.channel.QueueDeclare(
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

	err = b.channel.QueueBind(
		queue.Name, // queue name
		routeKey,   // routing key
		exchange,   // exchange
		false,
		nil,
	)

	return b
}
func (b *AMQP) Consume(cmd command.Handler, consumerName string, queueName string) {

	if b.conn == nil {
		b.connect()
	}

	ch, err := b.conn.Channel()

	if err != nil {
		log.Println("ERROR::: ", err.Error())
	}
	defer ch.Close()

	msgs, err := ch.Consume(
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
		for d := range msgs {
			cmd.Exec(d.RoutingKey, d.Body)
		}
	}()

	log.Printf(" [*] AMQP Ready")
	<-forever
}

func (b *AMQP) Dispatch(exchangeName string, routeKey string, message []byte) error {
	if b.conn == nil {
		b.connect()
	}

	return b.channel.Publish(
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

func (b *AMQP) Close() {
	b.channel.Close()
	b.conn.Close()
}

func (b *AMQP) connect() {

	var err error

	b.conn, err = amqp.Dial(b.dsn)

	if err != nil {
		log.Println("ERROR::: ", err)
	}

	b.channel, err = b.conn.Channel()

	if err != nil {
		log.Println("ERROR::: ", err)
	}

}
