package gateway

import (
	"chat_socket/model/gateway"
	"encoding/json"

	"github.com/gorilla/websocket"
)

func replyMessage(conn *websocket.Conn, binaryReq gateway.Req, resp []byte) {
	mReply := gateway.Resp{
		ReqIdentifier: binaryReq.ReqIdentifier,
		Data:          resp,
	}
	data, _ := json.Marshal(mReply)
	_ = conn.WriteMessage(websocket.BinaryMessage, data)
}
