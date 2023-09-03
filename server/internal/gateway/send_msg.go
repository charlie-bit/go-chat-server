package gateway

import (
	"chat_socket/server/config"
	"chat_socket/server/model/gateway"
	"chat_socket/server/pkg/proto/msg"
	"context"
	"github.com/charlie-bit/utils/basic_convert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

// use rpc_msg rpc client to send rpc_msg
func sendMsg(req gateway.Req) ([]byte, error) {
	msgClient, err := grpc.Dial(
		":"+basic_convert.NewBasicTypeConversion.Itoa(config.Cfg.MsgServerGrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := msg.NewMsgClient(msgClient)
	var msgReq = msg.SendMsgReq{
		SendID:  req.SendID,
		RecvID:  "charlie",
		Content: string(req.Data),
	}
	resp, err := client.SendMsg(context.Background(), &msgReq)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(resp)
}
