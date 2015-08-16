// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	MessageRequest
	MessageReply
*/
package message

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type MessageRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}

type MessageReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *MessageReply) Reset()         { *m = MessageReply{} }
func (m *MessageReply) String() string { return proto.CompactTextString(m) }
func (*MessageReply) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for MessageService service

type MessageServiceClient interface {
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (MessageService_SendMessageClient, error)
	SendMessageSimple(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (MessageService_SendMessageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MessageService_serviceDesc.Streams[0], c.cc, "/message.MessageService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceSendMessageClient{stream}
	return x, nil
}

type MessageService_SendMessageClient interface {
	Send(*MessageRequest) error
	Recv() (*MessageReply, error)
	grpc.ClientStream
}

type messageServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *messageServiceSendMessageClient) Send(m *MessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceSendMessageClient) Recv() (*MessageReply, error) {
	m := new(MessageReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) SendMessageSimple(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := grpc.Invoke(ctx, "/message.MessageService/SendMessageSimple", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageService service

type MessageServiceServer interface {
	SendMessage(MessageService_SendMessageServer) error
	SendMessageSimple(context.Context, *MessageRequest) (*MessageReply, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).SendMessage(&messageServiceSendMessageServer{stream})
}

type MessageService_SendMessageServer interface {
	Send(*MessageReply) error
	Recv() (*MessageRequest, error)
	grpc.ServerStream
}

type messageServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *messageServiceSendMessageServer) Send(m *MessageReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceSendMessageServer) Recv() (*MessageRequest, error) {
	m := new(MessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageService_SendMessageSimple_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(MessageRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(MessageServiceServer).SendMessageSimple(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessageSimple",
			Handler:    _MessageService_SendMessageSimple_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _MessageService_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
