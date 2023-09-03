package start_rpc

import (
	"github.com/charlie-bit/utils/basic_convert"
	"github.com/charlie-bit/utils/gzlog"
	"google.golang.org/grpc"
	"net"
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
	gzlog.Infof("start server failed, port : %d", rpcPort)
	return nil
}
