package main

import (
	"fmt"
	"strings"
	"sync"

	zmq "github.com/pebbe/zmq4"
)

type words struct {
  sync.Mutex
  found map[string]int
}

func main() {
	//  Socket to talk to server
	fmt.Println("Agrupando palavras vindas dos workers...")
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5556")

	var wg sync.WaitGroup

	//w := newWords()
	wg.Wait()

	fmt.Println("Ocorrências de cada palavra:")
	//mutex.Lock()
	// for word, count := range w.found {
	// 	if count > 1 {
	// 		fmt.Printf("%s: %d\n", word, count)
	// 	}
	// }
	//mutex.Unlock()
	m := newWords()

	for {
    m.Lock()

		msg, _ := subscriber.Recv(0)
		palavra := strings.Fields(msg)
		for _, word := range palavra {

      m.add(word, 1)

		}

		for pala, ocor := range m.found {
			if ocor > 1 {
				fmt.Printf("%s: %d\n", pala, ocor)
			}
    }
    m.Unlock()
	}

}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add (word string, n int){
  w.Lock()
  defer w.Unlock()

  count, ok := w.found[word]

  if !ok {
    w.found[word] = n
    return
  }

  w.found[word] = count + n
}