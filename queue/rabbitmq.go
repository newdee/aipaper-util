package queue

import (
	"fmt"
	"github.com/newdee/aipaper-util/config/business/common"
	"github.com/newdee/aipaper-util/log"
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

func reconnect() error {
	log.Debugf("MQ client is closed, trying to reconnect...")
	var err error
	conf, err := common.GetMsgQueueConfig()
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Vhost)
	conn, err = amqp091.Dial(uri)
	if err != nil {
		return fmt.Errorf("failed to reconnect to MQ: %v", err)
	}
	log.Infof("Successfully reconnected to MQ.")
	return nil
}

func SendMsg(queueName string, msg string) error {
	if conn == nil {
		return fmt.Errorf("MQ client not initialized")
	}
	// 检查连接是否已经关闭，如果已关闭则尝试进行重连
	if conn.IsClosed() {
		err := reconnect()
		if err != nil {
			return err
		}
	}
	// 打开通道
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()
	// 推送消息
	err = ch.Publish("", queueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})

	return err
}
