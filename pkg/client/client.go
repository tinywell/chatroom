package client

import (
	"context"
	"io"

	"google.golang.org/grpc"

	"github.com/tinywell/chatroom/pkg/proto"
)

// Client chat client
type Client struct {
	name   string
	reader io.Reader
	writer io.Writer
	chatC  proto.ChatRoomClient
}

// NewClient return a new Client
func NewClient(name string, conn *grpc.ClientConn, reader io.Reader, writer io.Writer) *Client {
	chatC := proto.NewChatRoomClient(conn)
	return &Client{
		name:   name,
		reader: reader,
		writer: writer,
		chatC:  chatC,
	}
}

// Start to chat
func (c *Client) Start() error {
	ctx := context.Background()
	msg := c.newSignMsg()
	stream, err := c.chatC.Chat(ctx)
	if err != nil {
		return err
	}
	stream.Send(msg)
	go c.sendMsg(stream)
	go c.recvMsg(stream)
	return nil
}

func (c *Client) sendMsg(stream proto.ChatRoom_ChatClient) {
	for {
		p := []byte{}
		_, err := c.reader.Read(p)
		if err != nil {
			return
		}
		msg := c.newChatMsg(string(p))
		err = stream.Send(msg)
		if err != nil {
			return
		}
	}
}

func (c *Client) recvMsg(stream proto.ChatRoom_ChatClient) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return
		}
		switch msg.Payload.(type) {
		case *proto.Message_Signin:
		case *proto.Message_Chatmsg:
			cmsg := msg.Payload.(*proto.Message_Chatmsg)
			_, err := c.writer.Write([]byte(cmsg.Chatmsg.Msg))
			if err != nil {
				return
			}
		}
	}
}

func (c *Client) stop() {

}

func (c *Client) newChatMsg(msg string) *proto.Message {
	cmsg := &proto.Message_Chatmsg{}
	cmsg.Chatmsg.Name = c.name
	cmsg.Chatmsg.Msg = msg
	return &proto.Message{Name: c.name, Payload: cmsg}
}

func (c *Client) newSignMsg() *proto.Message {
	smsg := &proto.Message_Signin{}
	smsg.Signin.Name = c.name
	return &proto.Message{Name: c.name, Payload: smsg}
}
