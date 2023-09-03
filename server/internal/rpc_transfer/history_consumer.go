package rpc_transfer

import (
	"chat_socket/server/internal/rpc_transfer/cache"
	relation "chat_socket/server/model/table"
	"chat_socket/server/pkg/constant"
	"chat_socket/server/pkg/kafka"
	"chat_socket/server/pkg/proto/msg"
	"context"
	"github.com/IBM/sarama"
	"github.com/charlie-bit/utils/third_party/go-redis"
	"google.golang.org/protobuf/proto"
	"sort"
	"strings"
)

type historyConsumer struct {
	historyConsumerGroup *kafka.MConsumerGroup
	rdb                  redis.UniversalClient
	convMInter           relation.ConversationModelInter
}

func (h historyConsumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h historyConsumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h historyConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for mvale := range claim.Messages() {
		h.handleMsg(mvale)
		session.MarkMessage(mvale, "consumed")
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
	convID := GetConversationIdByMsg(&content)
	// get cur seq
	_, err := cache.GetConvIDSeq(h.rdb, convID)
	if err != nil {
		return
	}
	// set conversation max seq
	convSeq, err := cache.SetConvIDSeq(h.rdb, convID, 1)
	if err != nil {
		return
	}
	// new conversation need create new record
	if convSeq == 1 {
		err = h.convMInter.Create(
			context.Background(), []*relation.ConversationModel{
				{
					ConversationType: constant.SuperGroupChatType,
					GroupID:          content.RecvID,
					OwnerUserID:      content.SendID,
					ConversationID:   convID,
				},
			},
		)
		if err != nil {
			return
		}
	}
	// storage message in mongo
	// mongoMsg := mongo_table.MsgDataModel{
	// 	SendID:      content.SendID,
	// 	RecvID:      content.RecvID,
	// 	GroupID:     content.RecvID,
	// 	ClientMsgID: content.RecvID,
	// 	ServerMsgID: content.SendID,
	// 	SessionType: content.SessionType,
	// 	Content:     content.Content,
	// 	Seq:         curSeq,
	// }
	// 返回值为true表示数据库存在该文档，false表示数据库不存在该文档
	// updateMsgModel := func(seq int64, i int) (bool, error) {
	// 	var (
	// 		res *mongo.UpdateResult
	// 		err error
	// 	)
	// 	docID := db.msg.GetDocID(conversationID, seq)
	// 	index := db.msg.GetMsgIndex(seq)
	// 	field := fields[i]
	// 	res, err = db.msgDocDatabase.UpdateMsg(ctx, docID, index, "msg", field)
	// 	if err != nil {
	// 		return false, err
	// 	}
	// 	return res.MatchedCount > 0, nil
	// }

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
