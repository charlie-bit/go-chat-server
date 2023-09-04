package gateway

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/charlie-bit/utils/basic_convert"
	"github.com/charlie-bit/utils/gzlog"
	"github.com/gorilla/websocket"
)

func WsRun(wsPort int) error {

	http.HandleFunc("/", wsHandler)
	return http.ListenAndServe(
		":"+basic_convert.NewBasicTypeConversion.Itoa(wsPort), nil,
	) // Start listening
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		HandshakeTimeout: time.Second * 30,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		gzlog.Errorf("build ws conn failed,err : %s", err.Error())
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("socket have panic err:", r, string(debug.Stack()))
		}
		_ = conn.Close()
	}()
	conn.SetReadLimit(51200)
	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	conn.SetPongHandler(
		func(_ string) error {
			return conn.SetReadDeadline(time.Now().Add(time.Second * 30))
		},
	)
	UserClientMap.Set(r.URL.Query().Get("userID"), conn)
	for {
		messageType, message, returnErr := conn.ReadMessage()
		if returnErr != nil {
			return
		}
		switch messageType {
		case websocket.BinaryMessage:
			_ = handleMsg(conn, message)
		case websocket.PingMessage:
		case websocket.CloseMessage:
			return
		}
	}
}
