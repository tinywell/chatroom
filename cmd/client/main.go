package main

import (
	"flag"
	"fmt"

	"github.com/tinywell/chatroom/pkg/client"
	"github.com/tinywell/chatroom/pkg/proto"
	"github.com/tinywell/chatroom/pkg/ui"

	"google.golang.org/grpc"
)

func main() {
	var serverAddr string
	var name string
	flag.StringVar(&serverAddr, "server", "localhost:8888", "server address")
	flag.StringVar(&name, "user", "test", "user name")
	flag.Parse()

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	c := client.NewClient(name, conn)
	c.Start()
	roomUI, err := ui.NewUI(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	inputC := roomUI.GetInputC()
	recv := c.GetRecvMsg()
	go func() {
		for input := range inputC {
			if len(input) == 0 {
				continue
			}
			c.SendMsg(input)
		}
	}()

	go func() {
		for msg := range recv {
			// fmt.Println(msg)
			if cmsg, ok := msg.GetPayload().(*proto.Message_Chatmsg); ok {
				roomUI.ApppendMsg(cmsg.Chatmsg.Msg, cmsg.Chatmsg.Name)
			}
		}
	}()

	roomUI.Open()

	// f := bufio.NewReader(os.Stdin)
	// for {
	// 	input, err := f.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	if len(input) == 1 { // empty line
	// 		continue
	// 	}
	// 	cmd := ""
	// 	fmt.Sscan(input, &cmd)
	// 	if cmd == "quit" {
	// 		break
	// 	}
	// 	c.SendMsg(input)
	// }
}

func getMsg(recv <-chan *proto.Message, roomUI *ui.UI) {

}
