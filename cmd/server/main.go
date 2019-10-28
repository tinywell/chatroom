package main

import (
	"net"

	"github.com/tinywell/chatroom/pkg/proto"
	"github.com/tinywell/chatroom/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := server.NewServer()
	proto.RegisterChatRoomServer(s, srv)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {

	}
	s.Serve(lis)
}
