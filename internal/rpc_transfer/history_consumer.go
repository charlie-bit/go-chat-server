package rpc_transfer

import (
	"chat_socket/internal/rpc_transfer/cache"
	unrelation2 "chat_socket/model/mongo_table"
	relation2 "chat_socket/model/table"
	"chat_socket/pkg/constant"
	kafka2 "chat_socket/pkg/kafka"
	"chat_socket/pkg/proto/msg"
	"context"
	"sort"
	"strings"

	"github.com/IBM/sarama"
	"github.com/charlie-bit/utils/third_party/go-redis"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/proto"
)

type historyConsumer struct {
	historyConsumerGroup *kafka2.MConsumerGroup
	rdb                  redis.UniversalClient
	convMInter           relation2.ConversationModelInter
	msgMongoInter        unrelation2.MsgDocModelInter
	pushProducer         *kafka2.Producer
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
	curSeq, err := cache.GetConvIDSeq(h.rdb, convID)
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
			context.Background(), []*relation2.ConversationModel{
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
	mongoMsg := unrelation2.MsgDataModel{
		SendID:      content.SendID,
		RecvID:      content.RecvID,
		GroupID:     content.RecvID,
		ClientMsgID: content.RecvID,
		ServerMsgID: content.SendID,
		SessionType: content.SessionType,
		Content:     content.Content,
		Seq:         curSeq,
	}
	// 返回值为true表示数据库存在该文档，false表示数据库不存在该文档
	updateMsgModel := func(seq int64, i int) (bool, error) {
		var (
			res *mongo.UpdateResult
			err error
		)
		docID := unrelation2.GetDocID(convID, mongoMsg.Seq)
		index := unrelation2.GetMsgIndex(seq)
		field := mongoMsg
		res, err = h.msgMongoInter.UpdateMsg(docID, index, "msg", field)
		if err != nil {
			return false, err
		}
		return res.MatchedCount > 0, nil
	}
	// 先查询mongo中是否存在该doc，如果有，优先更新
	// 如果没有，插入
	doc := unrelation2.MsgDocModel{
		DocID: unrelation2.GetDocID(convID, mongoMsg.Seq),
		Msg:   []*unrelation2.MsgInfoModel{{Msg: &mongoMsg}},
	}
	var tryUpdate bool
	if err := h.msgMongoInter.Create(&doc); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			tryUpdate = true // 以修改模式
		} else {
			return
		}
	}
	if tryUpdate {
		_, err = updateMsgModel(mongoMsg.Seq, int(unrelation2.GetMsgIndex(mongoMsg.Seq)))
		if err != nil {
			return
		}
	}
	// transfer message to push server
	_, _, _ = h.pushProducer.SendMsg(context.Background(), content.RecvID, string(message.Value))
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