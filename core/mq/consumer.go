package mq

import (
	"Blog/core/logs"

	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

type EmailMessageHandler struct{}

func (h *EmailMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	logs.Logs.Info("Process Send Email", zap.String("recevie", m.NSQDAddress), zap.String("Payload", string(m.Body)))

	return nil
}

func Consumer() error {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("email_topic", "email_channel", config)
	if err != nil {
		return err
	}

	// 添加Handler
	consumer.AddHandler(&EmailMessageHandler{})

	// 建立一个nsqd连接
	return consumer.ConnectToNSQD("192.168.433.178:4150")
}
