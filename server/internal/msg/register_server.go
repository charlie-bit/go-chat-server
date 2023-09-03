package msg

import (
	"chat_socket/server/config"
	"chat_socket/server/pkg/kafka"
	"chat_socket/server/pkg/proto/msg"
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
