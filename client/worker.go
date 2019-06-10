package main

import (
	zmq "github.com/pebbe/zmq4"
	"sync"

	"fmt"
)

func main() {
	w:=new(words)

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
		w.add(k)
		}

		//  Send results to sink
		sender.Send("", 0)
	}
}

var mutex sync.Mutex

type words struct {
	found map[string]int
}

func (w *words) add(word string) {
	mutex.Lock()
	defer mutex.Unlock()

	count, ok := w.found[word]
	if !ok {
		w.found[word] = 0
		return
	}
	w.found[word] = count + 1
}
/* Ex.: Outra forma de contar a ocorrência de palavras
func wordCount(str string) map[string]int {
    wordList := strings.Fields(str)
    counts := make(map[string]int)
    for _, word := range wordList {
        _, ok := counts[word]
        if ok {
            counts[word] += 1
        } else {
            counts[word] = 1
        }
    }
    return counts
}

func main() {
    strLine := "Australia Canada Germany Australia Japan Canada"
    for index,element := range wordCount(strLine){
        fmt.Println(index,"=>",element)
    }
}
}
*/
