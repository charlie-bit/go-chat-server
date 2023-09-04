package rpc_msg

import (
	"chat_socket/pkg/proto/msg"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/charlie-bit/utils/gmd5"
)

func (m MsgRPCServer) SendMsg(ctx context.Context, req *msg.SendMsgReq) (*msg.SendMsgResp, error) {
	var resp = &msg.SendMsgResp{}
	// validate message struct
	if req.Content == "" {
		return nil, fmt.Errorf("req message is nil")
	}
	// integrate rpc_msg basic info
	resp.ClientMsgID = getMsgID(req.RecvID)
	resp.ServerMsgID = getMsgID(req.SendID)
	resp.SendTime = time.Now().UnixNano() / 1e6
	// push rpc_msg to mq
	if _, _, err := m.producer.SendMsg(ctx, req.RecvID, req.Content); err != nil {
		return nil, err
	}
	return resp, nil
}

func getMsgID(sendID string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return gmd5.MD5(t + "-" + sendID + "-" + strconv.Itoa(rand.Int()))
}
