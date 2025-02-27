package queue

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/newdee/aipaper-util/config/business/common"
	"github.com/newdee/aipaper-util/log"
	"github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

var conn *amqp091.Connection
var channelPool chan *amqp091.Channel
var poolSize = 20
var mu sync.Mutex

func GetGlobalClient() *amqp091.Connection {
	return conn
}

func InitMQ(conf common.MsgQueueConfig) error {
	// 初始化MQ长连接
	var err error
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Vhost)
	fmt.Println("uri:", uri)
	// 提高MQ连接超时时间
	dial := amqp091.DefaultDial(60 * time.Second)
	conn, err = amqp091.DialConfig(uri, amqp091.Config{
		Heartbeat: 30 * time.Second, // 心跳间隔设置为30s，减少网络不稳时的误判
		Dial:      dial,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to MQ: %v", err)
	}

	// 初始化MQ的channel池
	channelPool = make(chan *amqp091.Channel, poolSize)
	for i := 0; i < poolSize; i++ {
		ch, err := conn.Channel()
		if err != nil {
			return fmt.Errorf("init mq channelPool failed, reason：[failed to open a channel: %v]", err)
		}
		channelPool <- ch
	}

	// 启动一个协程监听连接关闭事件
	go monitorConnection()

	return nil
}

func reconnect() error {
	log.Debugf("MQ client is closed, trying to reconnect...")
	mu.Lock()
	defer mu.Unlock()

	// 检查连接是否已经被其他协程重新建立
	if conn != nil && !conn.IsClosed() {
		return nil
	}

	// 从Apollo获取MQ配置
	conf, err := common.GetMsgQueueConfig()
	if err != nil {
		return fmt.Errorf("failed to get MQ config: %v", err)
	}

	// 重新建立长连接
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Vhost)
	// 提高MQ连接超时时间
	dial := amqp091.DefaultDial(60 * time.Second)
	conn, err = amqp091.DialConfig(uri, amqp091.Config{
		Heartbeat: 30 * time.Second, // 心跳间隔设置为30s，减少网络不稳时的误判
		Dial:      dial,
	})
	if err != nil {
		return fmt.Errorf("failed to reconnect to MQ: %v", err)
	}

	// 重新初始化channel池
	channelPool = make(chan *amqp091.Channel, poolSize)
	for i := 0; i < poolSize; i++ {
		ch, err := conn.Channel()
		if err != nil {
			return fmt.Errorf("failed to open a channel: %v", err)
		}
		channelPool <- ch
	}

	// 重新启动连接关闭监听
	go monitorConnection()

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

	// 从channel池中获取一个可用的channel
	var ch *amqp091.Channel
	select {
	case ch = <-channelPool:
		// 成功获取到一个 Channel
	case <-time.After(10 * time.Second): // 当channel池中暂时没有可用channel时，等待10s
		return fmt.Errorf("waiting to obtain an available channel timeout")
	}

	defer func() {
		// 将 Channel 归还到池中
		channelPool <- ch
	}()

	// 生成全局唯一的 message_id
	messageID := uuid.New().String()

	// 推送消息
	return ch.Publish("", queueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
		Headers:     amqp091.Table{"x-retry": 1}, // 初始化重试次数为1
		MessageId:   messageID,
	})
}

// 监听MQ连接是否关闭
func monitorConnection() {
	closeChan := conn.NotifyClose(make(chan *amqp091.Error))
	for err := range closeChan {
		log.ErrorfAlert("RabbitMQ connection closed unexpectedly: %v", err)
	}
}
