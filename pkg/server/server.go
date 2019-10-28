package server

import (
	"github.com/tinywell/chatroom/pkg/proto"
)

// Client is a client connectied to the chatroom
type Client struct {
	name   string
	stream proto.ChatRoom_ChatServer
}

// Server is the chatroom server
type Server struct {
}

// NewServer return a new Server
func NewServer() *Server {
	return nil
}

func (s *Server) start() {

}

// Chat is implement for proto.ChatRoomServer
func (s *Server) Chat(stream proto.ChatRoom_ChatServer) {

}
