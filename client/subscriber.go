package main

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
	"sync"
)

func main(){
	//  Socket to talk to server
	fmt.Println("Agrupando palavras vindas dos workers...")
	subscriber, _ := zmq.NewSocket(zmq.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5556")

	var wg sync.WaitGroup

	w := newWords()
	wg.Wait()

	fmt.Println("Ocorrências de cada palavra:")
	mutex.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	mutex.Unlock()
}
func newWords() *words {
	return &words{found: map[string]int{}}
}