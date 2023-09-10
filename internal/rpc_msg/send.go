package rpc_msg

import (
	"chat_socket/pkg/proto/msg"
	"context"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"strconv"
	"time"

	"github.com/charlie-bit/utils/gmd5"
)

func (m MsgRPCServer) SendMsg(ctx context.Context, req *msg.SendMsgReq) (*msg.SendMsgResp, error) {
	var resp = &msg.SendMsgResp{}
	// integrate rpc_msg basic info
	resp.ClientMsgID = getMsgID(req.Msg.RecvID)
	resp.ServerMsgID = getMsgID(req.Msg.SendID)
	resp.SendTime = time.Now().UnixNano() / 1e6
	// push rpc_msg to mq
	data, _ := proto.Marshal(req.Msg)
	if _, _, err := m.producer.SendMsg(ctx, req.Msg.GroupID, string(data)); err != nil {
		return nil, err
	}
	return resp, nil
}

func getMsgID(sendID string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return gmd5.MD5(t + "-" + sendID + "-" + strconv.Itoa(rand.Int()))
}
