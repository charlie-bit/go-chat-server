package rpc_push

import (
	"chat_socket/server/config"
	unrelation "chat_socket/server/model/mongo_table"
	relation "chat_socket/server/model/table"
	"chat_socket/server/pkg/kafka"
	"chat_socket/server/pkg/mongo"
	"chat_socket/server/pkg/mysql"
	"chat_socket/server/pkg/redis"

	"google.golang.org/grpc"
)

type PushRPCServer struct {
	hConsumer historyConsumer
}

func StartTransferServer(s *grpc.Server) {
	var srv PushRPCServer
	srv.hConsumer = historyConsumer{
		historyConsumerGroup: kafka.NewMConsumerGroup(
			[]string{config.Cfg.Kafka.MsgToPush.Topic},
			config.Cfg.Kafka.Addr,
			config.Cfg.Kafka.MsgToPush.Topic,
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
	mongodb, err := mongo.NewMongo()
	if err != nil {
		return
	}
	srv.hConsumer.convMInter = relation.NewConversationGorm(mdb)
	srv.hConsumer.msgMongoInter = unrelation.NewMsgMongoDriver(mongodb.GetDatabase())
	go srv.hConsumer.historyConsumerGroup.RegisterHandleAndConsumer(srv.hConsumer)
}
