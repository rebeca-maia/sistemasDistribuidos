package main

import (
	zmq "github.com/pebbe/zmq4"
	"sync"

	"fmt"
)

func main() {
	w:= new(words)

	//map de entrada
	m =make(map[string]int)

	//  Socket to receive messages on
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	//  Socket to send messages to
	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

	//  Process tasks forever
	for {
		s, _ := receiver.Recv(0)

		//  Simple progress indicator for the viewer
		fmt.Print(s + ".")

		//  Do the work
		for k:= range m{
		w.add(k,m[k])
		}

		//  Send results to sink
		sender.Send("", 0)
	}
}

var mutex sync.Mutex

type words struct {
	found map[string]int
}

func (w *words) add(word string, n int) {
	mutex.Lock()
	defer mutex.Unlock()

	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

