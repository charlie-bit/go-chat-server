package rpc_transfer

import (
	"chat_socket/server/config"
	relation "chat_socket/server/model/table"
	"chat_socket/server/pkg/kafka"
	"chat_socket/server/pkg/mysql"
	"chat_socket/server/pkg/redis"
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
	rdb, err := redis.NewRedis()
	if err != nil {
		return
	}
	srv.hConsumer.rdb = rdb
	mdb, err := mysql.NewMysqlGormDB()
	if err != nil {
		return
	}
	_ = mdb.AutoMigrate(&relation.ConversationModel{})
	srv.hConsumer.convMInter = relation.NewConversationGorm(mdb)
	go srv.hConsumer.historyConsumerGroup.RegisterHandleAndConsumer(srv.hConsumer)
}
