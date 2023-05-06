package main

import (
	"chat_socket/internal/tcp"
	"chat_socket/pkg/util"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":2023")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	util.SafeGO(func(args interface{}) {
		tcp.Broadcast()
	}, nil)

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		util.SafeGO(func(args interface{}) {
			tcp.HandleConn(args.(net.Conn))
		}, conn)
	}
}
