package main

import (
	"fmt"
	//"strings"
	"net"
	//"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"sync"
	//zmq "github.com/pebbe/zmq4"
)

// Words struct
type Words struct {
	Messagens string
}

//Menssage String
type Menssage Words

// Count func
func (m *Menssage) Count(word *Words, reply *int) error {
	fmt.Println("exc")
	fmt.Println(word.Messagens)
	return nil
}

func main() {

	menssage := new(Menssage)
	rpc.Register(menssage)
	/*	rpc.HandleHTTP() */

	/* err := http.ListenAndServe(":1243", nil)
	if err != nil {
		fmt.Println(err.Error())
	} */

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
	//  Socket to talk to server
	fmt.Println("Agrupando palavras vindas dos workers...")
	/* 	subscriber, _ := zmq.NewSocket(zmq.SUB)
	   	defer subscriber.Close()
	   	subscriber.Connect("tcp://localhost:5556") */

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
	//m := newWords()

	for {
		//m.Lock()
		//msg, _ := subscriber.Recv(0)
		//palavra := strings.Fields(msg)
		/* for _, word := range palavra {

		      m.add(word, 1)

				} */

		/* 		for pala, ocor := range m.found {
					if ocor > 1 {
						fmt.Printf("%s: %d\n", pala, ocor)
					}
		    } */
		//m.Unlock()
	}
}

/* func newWords() *Words {
	return &Words{found: map[string]int{}}
}

func (w *Words) add (word string, n int){
  //w.Lock()
  //defer w.Unlock()

  count, ok := w.found[word]

  if !ok {
    w.found[word] = n
    return
  }

  w.found[word] = count + n
} */

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
