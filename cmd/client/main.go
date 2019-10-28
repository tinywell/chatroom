package main

import (
	"fmt"

	"github.com/tinywell/chatroom/pkg/proto"

	"github.com/tinywell/chatroom/pkg/client"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	c := client.NewClient("test", conn)
	c.Start()
	recv := c.GetRecvMsg()
	go getMsg(recv)
	for {
		msg := ""
		_, err := fmt.Scanf("%s\n", &msg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.SendMsg(msg)
	}
}

func getMsg(recv <-chan *proto.Message) {
	for msg := range recv {
		fmt.Println(msg)
	}
}
