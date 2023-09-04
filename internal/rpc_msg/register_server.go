package rpc_msg

import (
	"chat_socket/config"
	"chat_socket/pkg/kafka"
	"chat_socket/pkg/proto/msg"

	"google.golang.org/grpc"
)

type MsgRPCServer struct {
	producer *kafka.Producer
}

func StartMsgServer(s *grpc.Server) {
	var srv MsgRPCServer
	srv.producer = kafka.NewKafkaProducer(config.Cfg.Kafka.LatestMsgToRedis.Topic, config.Cfg.Kafka.Addr)
	msg.RegisterMsgServer(s, srv)
}
