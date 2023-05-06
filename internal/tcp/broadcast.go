package tcp

import (
	"fmt"
	"strconv"
)

func Broadcast() {
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
