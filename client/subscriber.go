package main

import (
	"fmt"
	//"strings"
	//"net"
	//"net/http"
	//"net/rpc"
	//"net/rpc/jsonrpc"
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
	//zmq "github.com/pebbe/zmq4"
)

/* // Words struct
type Words struct {
	Messagens string
}

//Menssage String
type Menssage Words
*/
type Cont struct {
	sync.Mutex
	found map[string]int
}

var (
	j *Cont
)

func erroMensage(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var wg sync.WaitGroup
	/* 	menssage := new(Menssage)
	   	rpc.Register(menssage)

	   	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	   	checkError(err)

	   	listener, err := net.ListenTCP("tcp", tcpAddr)
	   	checkError(err) */

	//Rabbit
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	erroMensage(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	erroMensage(err, "Failed to open a channel")
	defer ch.Close()

	w := newWords()

	q, err := ch.QueueDeclare(
		"task_cont",
		true,
		false,
		false,
		false,
		nil,
	)
	erroMensage(err, "Falha ao declara fila")

	err = ch.Qos(1, 0, false)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	erroMensage(err, "Falha ao registra Consumidor")

	run := make(chan bool)
	go func() {
		for d := range msgs {
			wg.Add(1)
			palavra := string(d.Body)
			w.found[palavra]++
			wg.Done()
			d.Ack(false)
		}
		//imprime(w)

	}()
	wg.Wait()
	println("-------------------Contagem---------------------------")

	<-run
	print("Terminou")
	/* for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}

	//w := newWords()
	//  Socket to talk to server
	fmt.Println("Agrupando palavras vindas dos workers...")

	var wg sync.WaitGroup

	//w := newWords()
	wg.Wait()

	fmt.Println("Ocorrências de cada palavra:")
	j.Lock()
	for word, count := range j.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	j.Unlock()

	println("Hãm") */
	/* for {
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
		    }
		//m.Unlock()
	} */
}

func imprime(p *Cont) {
	for word, count := range p.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}

}

/* func addMap(p string) {
	pl := newWords()
	pl.Lock()
	pl.add(p, 0)
	pl.Unlock()

} */

func newWords() *Cont {
	return &Cont{found: map[string]int{}}
}

func (w *Cont) add(word string, n int) {
	w.Lock()
	defer w.Unlock()

	count, ok := w.found[word]

	if !ok {
		w.found[word] = n
		return
	}

	w.found[word] = count + n
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

/* //Count word function
func (m *Menssage) Count(word *Words, reply *int) error {
	fmt.Println(word.Messagens)
	pl := newWords()

	pl.add(word.Messagens, 1)

	return nil
} */
