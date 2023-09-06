package start_rpc

import (
	"net"

	"github.com/charlie-bit/utils/basic_convert"
	"google.golang.org/grpc"
)

func StartRPC(rpcPort int, rpcFn func(server *grpc.Server)) error {
	listener, err := net.Listen(
		"tcp", ":"+basic_convert.NewBasicTypeConversion.Itoa(rpcPort),
	)
	if err != nil {
		return err
	}
	defer listener.Close()
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(nil))
	if rpcFn != nil {
		rpcFn(server)
	}
	if err = server.Serve(listener); err != nil {
		return err
	}
	return nil
}
