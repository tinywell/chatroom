package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tinywell/chatroom/pkg/client"
	"github.com/tinywell/chatroom/pkg/proto"

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
	f := bufio.NewReader(os.Stdin)
	for {
		input, err := f.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(input) == 1 { // empty line
			continue
		}
		cmd := ""
		fmt.Sscan(input, &cmd)
		if cmd == "quit" {
			break
		}
		c.SendMsg(input)
	}
}

func getMsg(recv <-chan *proto.Message) {
	for msg := range recv {
		fmt.Println(msg)
	}
}
