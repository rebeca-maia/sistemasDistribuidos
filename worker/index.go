package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {

	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	//  Socket to send messages to
	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

	for {
		s, _ := receiver.Recv(0)

    //  Simple progress indicator for the viewer
		fmt.Println(s + " - ")
		sender.Send("", 0)
	}
}
