package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
)

type Producer struct {
	topic    string
	addr     []string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewKafkaProducer(topic string, addr []string) *Producer {
	p := Producer{}
	p.config = sarama.NewConfig()             // Instantiate a sarama Config
	p.config.Producer.Return.Successes = true // Whether to enable the successes channel to be notified after the message is sent successfully
	p.config.Producer.Return.Errors = true
	p.config.Producer.RequiredAcks = sarama.WaitForAll        // Set producer Message Reply level 0 1 all
	p.config.Producer.Partitioner = sarama.NewHashPartitioner // Set the hash-key automatic hash partition. When sending a message, you must specify the key value of the message. If there is no key, the partition will be selected randomly
	p.addr = addr
	p.topic = topic
	var producer sarama.SyncProducer
	var err error
	producer, err = sarama.NewSyncProducer(p.addr, p.config) // Initialize the client
	if err == nil {
		p.producer = producer
		return &p
	}
	p.producer = producer
	return &p
}

func (p *Producer) SendMsg(ctx context.Context, key string, msg string) (int32, int64, error) {
	kMsg := &sarama.ProducerMessage{}
	kMsg.Topic = p.topic
	kMsg.Key = sarama.StringEncoder(key)
	kMsg.Value = sarama.ByteEncoder(msg)
	if kMsg.Key.Length() == 0 || kMsg.Value.Length() == 0 {
		return 0, 0, fmt.Errorf("key or value is nil")
	}
	kMsg.Metadata = ctx
	return p.producer.SendMessage(kMsg)
}
