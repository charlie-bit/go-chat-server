package rpc_push

import (
	"chat_socket/config"
	"chat_socket/pkg/kafka"
	"chat_socket/pkg/proto/gateway"
	"chat_socket/pkg/proto/msg"
	"context"
	"log"

	"github.com/IBM/sarama"
	"github.com/charlie-bit/utils/basic_convert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type historyConsumer struct {
	historyConsumerGroup *kafka.MConsumerGroup
}

func (h historyConsumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h historyConsumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h historyConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		h.handleMsg(message.Value)
		session.MarkMessage(message, "")
	}
	return nil
}

func (h historyConsumer) handleMsg(message []byte) {
	var content msg.SendMsgReq
	if err := proto.Unmarshal(message, &content); err != nil {
		return
	}
	// TODO get user member id from group

	// get online connection from gateway
	// online push message
	conn, err := grpc.Dial(
		":"+basic_convert.NewBasicTypeConversion.Itoa(config.Cfg.GatewayGrpcPort), grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	msgClient := gateway.NewGatewayClient(conn)
	_, _ = msgClient.SuperGroupOnlineBatchPushOneMsg(
		context.Background(),
		&gateway.OnlineBatchPushOneMsgReq{MsgData: &content, PushToUserIDs: []string{content.RecvID}},
	)
}
