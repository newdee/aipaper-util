package queue

import (
	"fmt"
	"github.com/newdee/aipaper-util/config/business/common"
	"github.com/rabbitmq/amqp091-go"
)

var conn *amqp091.Connection

func GetGlobalClient() *amqp091.Connection {
	return conn
}

func InitMQ(conf common.MsgQueueConfig) error {
	var err error
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Vhost)
	fmt.Println("uri:", uri)
	conn, err = amqp091.Dial(uri)
	return err
}

func SendMsg(queueName string, msg string) error {
	if conn == nil {
		return fmt.Errorf("MQ client not initialized")
	}
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	// queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	// if err != nil {
	// 	return fmt.Errorf("failed to declare a queue: %v", err)
	// }

	err = ch.Publish("", queueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})

	return err
}
