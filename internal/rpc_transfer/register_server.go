package rpc_transfer

import (
	"chat_socket/config"
	"chat_socket/model/mongo_table"
	relation2 "chat_socket/model/table"
	kafka2 "chat_socket/pkg/kafka"
	"chat_socket/pkg/mongo"
	"chat_socket/pkg/mysql"
	"chat_socket/pkg/redis"

	"google.golang.org/grpc"
)

type TransferRPCServer struct {
	hConsumer historyConsumer
}

func StartTransferServer(s *grpc.Server) {
	var srv TransferRPCServer
	srv.hConsumer = historyConsumer{
		historyConsumerGroup: kafka2.NewMConsumerGroup(
			[]string{config.Cfg.Kafka.LatestMsgToRedis.Topic},
			config.Cfg.Kafka.Addr,
			config.Cfg.Kafka.LatestMsgToRedis.Topic,
		),
	}
	srv.hConsumer.pushProducer = kafka2.NewKafkaProducer(config.Cfg.Kafka.MsgToPush.Topic, config.Cfg.Kafka.Addr)
	rdb, err := redis.NewRedis()
	if err != nil {
		return
	}
	srv.hConsumer.rdb = rdb
	mdb, err := mysql.NewMysqlGormDB()
	if err != nil {
		return
	}
	err = mdb.AutoMigrate(&relation2.ConversationModel{})
	if err != nil {
		return
	}
	mongodb, err := mongo.NewMongo()
	if err != nil {
		return
	}
	srv.hConsumer.convMInter = relation2.NewConversationGorm(mdb)
	srv.hConsumer.msgMongoInter = unrelation.NewMsgMongoDriver(mongodb.GetDatabase())
	go srv.hConsumer.historyConsumerGroup.RegisterHandleAndConsumer(srv.hConsumer)
}
