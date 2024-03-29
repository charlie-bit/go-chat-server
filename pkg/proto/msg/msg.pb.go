// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.3
// source: msg/msg.proto

package msg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *MsgData `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMsgReq) Reset() {
	*x = SendMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReq) ProtoMessage() {}

func (x *SendMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReq.ProtoReflect.Descriptor instead.
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{0}
}

func (x *SendMsgReq) GetMsg() *MsgData {
	if x != nil {
		return x.Msg
	}
	return nil
}

type SendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerMsgID string `protobuf:"bytes,1,opt,name=serverMsgID,proto3" json:"serverMsgID,omitempty"` // 服务端消息id
	ClientMsgID string `protobuf:"bytes,2,opt,name=clientMsgID,proto3" json:"clientMsgID,omitempty"` // 客户端消息id
	SendTime    int64  `protobuf:"varint,3,opt,name=sendTime,proto3" json:"sendTime,omitempty"`      // 消息时间
}

func (x *SendMsgResp) Reset() {
	*x = SendMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgResp) ProtoMessage() {}

func (x *SendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgResp.ProtoReflect.Descriptor instead.
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{1}
}

func (x *SendMsgResp) GetServerMsgID() string {
	if x != nil {
		return x.ServerMsgID
	}
	return ""
}

func (x *SendMsgResp) GetClientMsgID() string {
	if x != nil {
		return x.ClientMsgID
	}
	return ""
}

func (x *SendMsgResp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

type MsgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendID         string `protobuf:"bytes,1,opt,name=sendID,proto3" json:"sendID,omitempty"`
	RecvID         string `protobuf:"bytes,2,opt,name=recvID,proto3" json:"recvID,omitempty"`
	GroupID        string `protobuf:"bytes,3,opt,name=groupID,proto3" json:"groupID,omitempty"`
	ClientMsgID    string `protobuf:"bytes,4,opt,name=clientMsgID,proto3" json:"clientMsgID,omitempty"`
	ServerMsgID    string `protobuf:"bytes,5,opt,name=serverMsgID,proto3" json:"serverMsgID,omitempty"`
	SenderNickname string `protobuf:"bytes,7,opt,name=senderNickname,proto3" json:"senderNickname,omitempty"`
	SenderFaceURL  string `protobuf:"bytes,8,opt,name=senderFaceURL,proto3" json:"senderFaceURL,omitempty"`
	SessionType    int32  `protobuf:"varint,9,opt,name=sessionType,proto3" json:"sessionType,omitempty"`
	MsgFrom        int32  `protobuf:"varint,10,opt,name=msgFrom,proto3" json:"msgFrom,omitempty"`
	ContentType    int32  `protobuf:"varint,11,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Content        []byte `protobuf:"bytes,12,opt,name=content,proto3" json:"content,omitempty"`
	Seq            int64  `protobuf:"varint,14,opt,name=seq,proto3" json:"seq,omitempty"`
	SendTime       int64  `protobuf:"varint,15,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
	CreateTime     int64  `protobuf:"varint,16,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Status         int32  `protobuf:"varint,17,opt,name=status,proto3" json:"status,omitempty"`
	IsRead         bool   `protobuf:"varint,18,opt,name=isRead,proto3" json:"isRead,omitempty"`
	Ex             string `protobuf:"bytes,23,opt,name=ex,proto3" json:"ex,omitempty"`
}

func (x *MsgData) Reset() {
	*x = MsgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgData) ProtoMessage() {}

func (x *MsgData) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgData.ProtoReflect.Descriptor instead.
func (*MsgData) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{2}
}

func (x *MsgData) GetSendID() string {
	if x != nil {
		return x.SendID
	}
	return ""
}

func (x *MsgData) GetRecvID() string {
	if x != nil {
		return x.RecvID
	}
	return ""
}

func (x *MsgData) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *MsgData) GetClientMsgID() string {
	if x != nil {
		return x.ClientMsgID
	}
	return ""
}

func (x *MsgData) GetServerMsgID() string {
	if x != nil {
		return x.ServerMsgID
	}
	return ""
}

func (x *MsgData) GetSenderNickname() string {
	if x != nil {
		return x.SenderNickname
	}
	return ""
}

func (x *MsgData) GetSenderFaceURL() string {
	if x != nil {
		return x.SenderFaceURL
	}
	return ""
}

func (x *MsgData) GetSessionType() int32 {
	if x != nil {
		return x.SessionType
	}
	return 0
}

func (x *MsgData) GetMsgFrom() int32 {
	if x != nil {
		return x.MsgFrom
	}
	return 0
}

func (x *MsgData) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *MsgData) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *MsgData) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *MsgData) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *MsgData) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *MsgData) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MsgData) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *MsgData) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

var File_msg_msg_proto protoreflect.FileDescriptor

var file_msg_msg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x2c, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x71, 0x12, 0x1e, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x6d, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0xeb, 0x03, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x61, 0x63,
	0x65, 0x55, 0x52, 0x4c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x73, 0x67, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73,
	0x67, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x73, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61,
	0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x32,
	0x33, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73,
	0x67, 0x12, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x10, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x1b, 0x5a, 0x19, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x73,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_msg_proto_rawDescOnce sync.Once
	file_msg_msg_proto_rawDescData = file_msg_msg_proto_rawDesc
)

func file_msg_msg_proto_rawDescGZIP() []byte {
	file_msg_msg_proto_rawDescOnce.Do(func() {
		file_msg_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_msg_proto_rawDescData)
	})
	return file_msg_msg_proto_rawDescData
}

var file_msg_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_msg_msg_proto_goTypes = []interface{}{
	(*SendMsgReq)(nil),  // 0: msg.SendMsgReq
	(*SendMsgResp)(nil), // 1: msg.SendMsgResp
	(*MsgData)(nil),     // 2: msg.MsgData
}
var file_msg_msg_proto_depIdxs = []int32{
	2, // 0: msg.SendMsgReq.msg:type_name -> msg.MsgData
	0, // 1: msg.msg.SendMsg:input_type -> msg.SendMsgReq
	1, // 2: msg.msg.SendMsg:output_type -> msg.SendMsgResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_msg_msg_proto_init() }
func file_msg_msg_proto_init() {
	if File_msg_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_msg_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_msg_proto_goTypes,
		DependencyIndexes: file_msg_msg_proto_depIdxs,
		MessageInfos:      file_msg_msg_proto_msgTypes,
	}.Build()
	File_msg_msg_proto = out.File
	file_msg_msg_proto_rawDesc = nil
	file_msg_msg_proto_goTypes = nil
	file_msg_msg_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// 发送消息
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := c.cc.Invoke(ctx, "/msg.msg/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// 发送消息
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}

func RegisterMsgServer(s *grpc.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.msg/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msg.msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _Msg_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg/msg.proto",
}
