package gateway

import (
	"chat_socket/pkg/proto/gateway"
	"google.golang.org/grpc"
)

type GatewayRPCServer struct {
}

func StartGatewayRPCServer(s *grpc.Server) {
	var srv GatewayRPCServer
	gateway.RegisterGatewayServer(s, srv)
}
