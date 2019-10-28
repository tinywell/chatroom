// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChatRoom service

type ChatRoomClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (ChatRoom_ChatClient, error)
}

type chatRoomClient struct {
	cc *grpc.ClientConn
}

func NewChatRoomClient(cc *grpc.ClientConn) ChatRoomClient {
	return &chatRoomClient{cc}
}

func (c *chatRoomClient) Chat(ctx context.Context, opts ...grpc.CallOption) (ChatRoom_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChatRoom_serviceDesc.Streams[0], c.cc, "/proto.ChatRoom/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRoomChatClient{stream}
	return x, nil
}

type ChatRoom_ChatClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatRoomChatClient struct {
	grpc.ClientStream
}

func (x *chatRoomChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatRoomChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChatRoom service

type ChatRoomServer interface {
	Chat(ChatRoom_ChatServer) error
}

func RegisterChatRoomServer(s *grpc.Server, srv ChatRoomServer) {
	s.RegisterService(&_ChatRoom_serviceDesc, srv)
}

func _ChatRoom_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatRoomServer).Chat(&chatRoomChatServer{stream})
}

type ChatRoom_ChatServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatRoomChatServer struct {
	grpc.ServerStream
}

func (x *chatRoomChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatRoomChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatRoom_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChatRoom",
	HandlerType: (*ChatRoomServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _ChatRoom_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}

func init() { proto1.RegisterFile("server.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x3c, 0xc9, 0xf9,
	0xb9, 0xb9, 0xf9, 0x79, 0x10, 0x41, 0x23, 0x0b, 0x2e, 0x0e, 0xe7, 0x8c, 0xc4, 0x92, 0xa0, 0xfc,
	0xfc, 0x5c, 0x21, 0x1d, 0x2e, 0x16, 0x10, 0x5b, 0x88, 0x0f, 0x22, 0xa7, 0xe7, 0x9b, 0x5a, 0x5c,
	0x9c, 0x98, 0x9e, 0x2a, 0x85, 0xc6, 0x57, 0x62, 0xd0, 0x60, 0x34, 0x60, 0x4c, 0x62, 0x03, 0x0b,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x14, 0x60, 0x8e, 0x65, 0x00, 0x00, 0x00,
}