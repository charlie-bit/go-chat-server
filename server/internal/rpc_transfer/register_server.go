package rpc_transfer

import (
	"chat_socket/server/config"
	"chat_socket/server/pkg/kafka"
	"google.golang.org/grpc"
)

type TransferRPCServer struct {
	producer  *kafka.Producer
	hConsumer historyConsumer
}

func StartTransferServer(s *grpc.Server) {
	var srv TransferRPCServer
	srv.producer = kafka.NewKafkaProducer(config.Cfg.Kafka.LatestMsgToRedis.Topic, config.Cfg.Kafka.Addr)
	srv.hConsumer = historyConsumer{
		historyConsumerGroup: kafka.NewMConsumerGroup(
			[]string{config.Cfg.Kafka.LatestMsgToRedis.Topic},
			config.Cfg.Kafka.Addr,
			config.Cfg.Kafka.LatestMsgToRedis.Topic,
		),
	}
	go srv.hConsumer.historyConsumerGroup.RegisterHandleAndConsumer(srv.hConsumer)
}
