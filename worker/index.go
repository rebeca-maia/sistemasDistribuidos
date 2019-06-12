package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"strconv"
	"strings"
	"time"
)

//Words struct
type Words struct {
	//messagens strinc slice
	Messagens string
}

func main() {

	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("erro ao Conectar", err)
	}

	m := make([]string, 2)

	for {
		s, _ := receiver.Recv(0)
		m = strings.Split(s, " ")

		for _, w := range m {
			println(w)
			var reply int
      Msg := &Words{w}
      print("chamando")
			err = client.Call("Menssage.Count", Msg, &reply)
			if err != nil {
				log.Fatal("erro ao Chama ", err)
			}
		}

		fmt.Println(m, " - ")

		//var reply []string

		//err = client.Call("Menssage.Count", msg, &reply)

		msec, _ := strconv.Atoi(s)
		time.Sleep(time.Duration(msec) * time.Millisecond)
		//  Simple progress indicator for the viewer
		sender.Send("", 0)
		//msg := fmt.Sprintf("%s", m)
		// publisher.Send(msg, 0)
	}

}
