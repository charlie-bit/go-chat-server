package main

import (
	"fmt"
	"github.com/charlie-bit/utils/safe_goroutine"
	"net"
	"strconv"
)

func main() {
	// 建立监听
	listen, err := net.Listen("tcp", ":2023")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	// 监听客户端链接
	safe_goroutine.SafeGo(func() {
		broadcast()
	})

	// 监听客户端消息
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		HandleConn(conn)
	}
}

func broadcast() {
	users := make(map[*User]struct{})

	for {
		select {
		case user := <-enterChannel:
			users[user] = struct{}{}

			// 进入聊天室，给当前所有用户发送欢迎信息
			MessageChannel <- "welcome, " + strconv.Itoa(user.ID)
		case user := <-leaveChannel:
			// 用户离开
			delete(users, user)
		case msg := <-MessageChannel:
			// 向非自己的所有用户广播消息
			fmt.Println(msg)
		}
	}
}
