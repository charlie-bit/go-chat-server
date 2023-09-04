package main

import (
	"chat_socket/server/config"
	"chat_socket/server/internal/rpc_transfer"
	"chat_socket/server/pkg/start_rpc"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/charlie-bit/utils/gzlog"
)

func main() {
	var (
		configFile  string
		ping        bool
		showVersion bool
		grpcPort    int
	)

	flag.StringVar(&configFile, "f", "server/config/setting.yaml", "config file")
	flag.BoolVar(&ping, "ping", false, "Ping server.")
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.IntVar(&grpcPort, "grpc_port", 0, "grpc port")
	flag.Parse()

	// get config
	err := config.InitConfig(configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	gzlog.Init()
	defer func() {
		gzlog.Exit()
	}()

	if err = run(grpcPort); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}

func run(port int) error {
	if port == 0 {
		port = config.Cfg.TransferServerGrpcPort
	}
	return start_rpc.StartRPC(port, rpc_transfer.StartTransferServer)
}
