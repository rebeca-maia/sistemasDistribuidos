package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"strconv"
	"strings"
	"time"
)

//var m := []string

func main() {

	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	//  Socket to send messages to
	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*5556")

	m := make([]string, 2)

	for {
		s, _ := receiver.Recv(0)
		m = strings.Split(s, " ")

		//fmt.Println("m na posição 1:  -  ", m[1])

		fmt.Println(m, " - ")

		msec, _ := strconv.Atoi(s)
		time.Sleep(time.Duration(msec) * time.Millisecond)
		//  Simple progress indicator for the viewer
		fmt.Println(s)
		sender.Send("", 0)
		msg := fmt.Sprintf("%s", m)
		publisher.Send(msg, 0)
	}

}
