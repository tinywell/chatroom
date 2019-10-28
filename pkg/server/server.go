package server

import (
	"sync"

	"github.com/tinywell/chatroom/pkg/proto"
)

// Client is a client connectied to the chatroom
type Client struct {
	name   string
	stream proto.ChatRoom_ChatServer
}

// Server is the chatroom server
type Server struct {
	clients map[string]*Client
	mutex   sync.RWMutex
}

// NewServer return a new Server
func NewServer() *Server {
	return &Server{
		clients: make(map[string]*Client),
	}
}

// Chat is implement for proto.ChatRoomServer
func (s *Server) Chat(stream proto.ChatRoom_ChatServer) {
	for {
		msg := &proto.Message{}
		err := stream.RecvMsg(msg)
		if err != nil {
			return
		}
		switch msg.Payload.(type) {
		case *proto.Message_Chatmsg:
			s.broadCast(msg)
		case *proto.Message_Signin:
			m := msg.Payload.(*proto.Message_Signin)
			s.newClient(m.Signin.Name, stream)
		default:
			return
		}
	}
}

func (s *Server) newClient(name string, stream proto.ChatRoom_ChatServer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	client := &Client{
		name:   name,
		stream: stream,
	}
	s.clients[name] = client
}
func (s *Server) broadCast(msg *proto.Message) {
	for _, c := range s.clients {
		err := c.stream.SendMsg(msg)
		if err != nil {
			continue
		}
	}
}
