package main

import (
	"fmt"
	"net"
	"time"
)

type User struct {
	ID      int       `json:"id"`
	EnterAt time.Time `json:"enter_at"`
}

func main() {
	// 建立监听
	listen, err := net.Listen("tcp", ":2023")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	// 监听客户端消息
	var (
		enterChannel = make(chan *User, 100)
	)

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println(conn.(*net.TCPConn).SetKeepAlive(true))
		fmt.Println(conn.(*net.TCPConn).SetKeepAlivePeriod(time.Second * 10))

		for {
			var message = make([]byte, 1024)
			n, err := conn.Read(message)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(string(message), n, err)

			var user = &User{
				EnterAt: time.Now(),
			}

			enterChannel <- user
			conn.Close()
		}
	}
}
