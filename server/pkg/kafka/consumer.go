package kafka

import (
	"context"
	"github.com/IBM/sarama"
)

type MConsumerGroup struct {
	sarama.ConsumerGroup
	group  string
	topics []string
}

func NewMConsumerGroup(topics, addr []string, group string) *MConsumerGroup {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Return.Errors = false
	consumerGroup, err := sarama.NewConsumerGroup(addr, group, config)
	if err != nil {
		panic(err.Error())
	}
	return &MConsumerGroup{
		consumerGroup,
		group,
		topics,
	}
}

func (mc *MConsumerGroup) RegisterHandleAndConsumer(handler sarama.ConsumerGroupHandler) {
	ctx := context.Background()
	for {
		err := mc.ConsumerGroup.Consume(ctx, mc.topics, handler)
		if err != nil {
			panic(err.Error())
		}
	}
}
