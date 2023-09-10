package gateway

import (
	gm "chat_socket/model/gateway"
	"chat_socket/pkg/constant"
	"encoding/json"
	"github.com/charlie-bit/utils/gzlog"

	"github.com/gorilla/websocket"
)

func handleMsg(conn *websocket.Conn, message []byte) error {
	// unmarshal request message byte
	var gmReq gm.Req
	if err := json.Unmarshal(message, &gmReq); err != nil {
		return err
	}
	// handle message with specific identifier
	var (
		resp []byte
		err  error
	)
	switch gmReq.ReqIdentifier {
	case constant.WSSendMsg:
		resp, err = sendMsg(gmReq)
		if err != nil {
			gzlog.Errorf("send message failed, err : %s", err.Error())
			return err
		}
	}
	replyMessage(conn, gmReq, resp)
	return nil
}
