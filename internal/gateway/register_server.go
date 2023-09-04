package gateway

import (
	"chat_socket/pkg/proto/gateway"
	"context"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type GatewayRPCServer struct {
}

func (g GatewayRPCServer) SuperGroupOnlineBatchPushOneMsg(
	ctx context.Context, req *gateway.OnlineBatchPushOneMsgReq,
) (*gateway.OnlineBatchPushOneMsgResp, error) {
	var resp = &gateway.OnlineBatchPushOneMsgResp{}
	resp.Resp = make([]*gateway.SingleMsgToUserPlatform, 0, len(req.PushToUserIDs))
	for _, id := range req.PushToUserIDs {
		conns, _, _ := UserClientMap.Get(id)
		for _, conn := range conns {
			if conn != nil {
				data, _ := proto.Marshal(req.MsgData)
				err := conn.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					return nil, err
				}
				resp.Resp = append(
					resp.Resp, &gateway.SingleMsgToUserPlatform{
						RecvID: id,
					},
				)
			}
		}
	}
	return resp, nil
}

func StartGatewayRPCServer(s *grpc.Server) {
	var srv GatewayRPCServer
	gateway.RegisterGatewayServer(s, srv)
}
