package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/tinywell/chatroom/pkg/proto"
	"github.com/tinywell/chatroom/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8888, "listen port")
	flag.Parse()

	s := grpc.NewServer()
	srv := server.NewServer()
	proto.RegisterChatRoomServer(s, srv)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	s.Serve(lis)
}
