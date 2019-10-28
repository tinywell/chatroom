package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/tinywell/chatroom/pkg/proto"
)

// Client chat client
type Client struct {
	name string
	// chatC  proto.ChatRoomClient
	stream proto.ChatRoom_ChatClient
	recvC  chan *proto.Message
}

// NewClient return a new Client
func NewClient(name string, conn *grpc.ClientConn) *Client {
	c := &Client{
		name:  name,
		recvC: make(chan *proto.Message),
	}

	chatC := proto.NewChatRoomClient(conn)
	ctx := context.Background()
	msg := c.newSignMsg()
	stream, err := chatC.Chat(ctx)
	if err != nil {
		return nil
	}
	stream.Send(msg)
	c.stream = stream
	return c
}

// Start to chat
func (c *Client) Start() error {
	// go c.sendMsg(stream)
	go c.recvMsg()
	return nil
}

// SendMsg is used to send msg to chatroom server
func (c *Client) SendMsg(msg string) {
	cm := c.newChatMsg(msg)
	err := c.stream.Send(cm)
	if err != nil {
		fmt.Println(err)
	}
}

// GetRecvMsg return receive msg channel
func (c *Client) GetRecvMsg() <-chan *proto.Message {
	return c.recvC
}

func (c *Client) recvMsg() {
	for {
		msg, err := c.stream.Recv()
		if err != nil {
			return
		}
		c.recvC <- msg
		// switch msg.Payload.(type) {
		// case *proto.Message_Signin:
		// case *proto.Message_Chatmsg:
		// 	cmsg := msg.Payload.(*proto.Message_Chatmsg)
		// 	c.recvC <- cmsg.Chatmsg
		// }
	}
}

func (c *Client) stop() {

}

func (c *Client) newChatMsg(msg string) *proto.Message {
	cmsg := &proto.Message_Chatmsg{Chatmsg: &proto.ChatMsg{}}
	cmsg.Chatmsg.Name = c.name
	cmsg.Chatmsg.Msg = msg
	return &proto.Message{Name: c.name, Payload: cmsg}
}

func (c *Client) newSignMsg() *proto.Message {
	smsg := &proto.Message_Signin{Signin: &proto.SignIn{Name: c.name}}
	return &proto.Message{Name: c.name, Payload: smsg}
}
