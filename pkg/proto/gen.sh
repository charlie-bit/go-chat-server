protoc --go_out=plugins=grpc:./msg --go_opt=module=chat_socket/pkg/proto/msg msg/msg.proto
protoc --go_out=plugins=grpc:./gateway --go_opt=module=chat_socket/pkg/proto/gateway gateway/gateway.proto