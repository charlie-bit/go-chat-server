package mongo_table

type MsgDataModel struct {
	SendID           string          `bson:"send_id"`
	RecvID           string          `bson:"recv_id"`
	GroupID          string          `bson:"group_id"`
	ClientMsgID      string          `bson:"client_msg_id"`
	ServerMsgID      string          `bson:"server_msg_id"`
	SenderPlatformID int32           `bson:"sender_platform_id"`
	SenderNickname   string          `bson:"sender_nickname"`
	SenderFaceURL    string          `bson:"sender_face_url"`
	SessionType      int32           `bson:"session_type"`
	MsgFrom          int32           `bson:"msg_from"`
	ContentType      int32           `bson:"content_type"`
	Content          string          `bson:"content"`
	Seq              int64           `bson:"seq"`
	SendTime         int64           `bson:"send_time"`
	CreateTime       int64           `bson:"create_time"`
	Status           int32           `bson:"status"`
	IsRead           bool            `bson:"is_read"`
	Options          map[string]bool `bson:"options"`
	AtUserIDList     []string        `bson:"at_user_id_list"`
	AttachedInfo     string          `bson:"attached_info"`
	Ex               string          `bson:"ex"`
}
