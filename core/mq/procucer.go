package mq

import "github.com/nsqio/go-nsq"

type NsqProducer struct {
	Addr  string
	Topic string
	Msg   []byte
}

func (p *NsqProducer) Producer() error {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(p.Addr, config)
	if err != nil {
		return err
	}
	// producer.PublishAsync()
	// 发送消息
	return producer.Publish(p.Topic, p.Msg)
}
