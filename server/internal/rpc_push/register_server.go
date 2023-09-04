package rpc_push

import (
	"chat_socket/server/config"
	"chat_socket/server/pkg/kafka"

	"google.golang.org/grpc"
)

type PushRPCServer struct {
	hConsumer historyConsumer
}

func StartPushServer(s *grpc.Server) {
	var srv PushRPCServer
	srv.hConsumer = historyConsumer{
		historyConsumerGroup: kafka.NewMConsumerGroup(
			[]string{config.Cfg.Kafka.MsgToPush.Topic},
			config.Cfg.Kafka.Addr,
			config.Cfg.Kafka.MsgToPush.Topic,
		),
	}
	go srv.hConsumer.historyConsumerGroup.RegisterHandleAndConsumer(srv.hConsumer)
}
