package rpc_transfer

import (
	"chat_socket/server/pkg/constant"
	"chat_socket/server/pkg/kafka"
	"chat_socket/server/pkg/proto/msg"
	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
	"sort"
	"strings"
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
	for msg := range claim.Messages() {
		h.handleMsg(msg)
		session.MarkMessage(msg, "consumed")
	}
	return nil
}

func (h historyConsumer) handleMsg(message *sarama.ConsumerMessage) {
	// unmarshal msg
	var content msg.SendMsgReq
	if err := proto.Unmarshal(message.Value, &content); err != nil {
		return
	}
	// get conversation id by msg
	// convID := GetConversationIdByMsg(&content)

	// set conversation max seq

	// new conversation need create new record

	//

	// storage message

	// transfer message to push server
}

func GetConversationIdByMsg(req *msg.SendMsgReq) string {
	l := []string{req.RecvID}
	sort.Strings(l)
	switch req.SessionType {
	case constant.SuperGroupChatType:
		return "sg" + strings.Join(l, "_")
	}
	return ""
}
