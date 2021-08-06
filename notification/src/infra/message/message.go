package message

import "sword-health/notification/application/command"

type Broker interface {
	QueueDeclare(exchange string, queueName string, routeKey string) Broker
	Consume(cmd command.Handler, consumerName string, queueName string)
	Dispatch(exchangeName string, routeKey string, message []byte) error
}
