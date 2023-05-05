package rabbitMQ

import (
	internal "github.com/SafetyLink/messageService/internal/infra"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func NewRabbitMQProvider(logger *zap.Logger, config *internal.Config) *amqp.Connection {
	conn, err := amqp.Dial(config.RabbitMQ.ConnectionUrl)
	if err != nil {
		logger.Panic("failed to connect to RabbitMQ")
	}

	logger.Info("connected to rabbit mq")
	return conn
}
