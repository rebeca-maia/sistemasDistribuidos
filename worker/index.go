package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
	"log"
	//"net/rpc/jsonrpc"
	"strconv"
	"strings"
  "time"
  "github.com/streadway/amqp"
)

//Words struct
type Words struct {
	//messagens strinc slice
	Messagens string
}

func erroMensage(err error, msg string){
  if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

  // Rabbit
  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	erroMensage(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	erroMensage(err, "Failed to open a channel")
	defer ch.Close()

  q, err := ch.QueueDeclare(
    "task_cont",
    true,
    false,
    false,
    false,
    nil,
  )

  erroMensage(err, "Falha ao declara a Fila de tarefas")

/* 	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("erro ao Conectar", err)
	} */

	m := make([]string, 2)

	for {
		s, _ := receiver.Recv(0)
		m = strings.Split(s, " ")

		for _, w := range m {
			println(w)
			/* var reply int
      Msg := &Words{w}
      print("chamando : ")
			err = client.Call("Menssage.Count", Msg, &reply)
			if err != nil {
				log.Fatal("erro ao Chama ", err)
      } */

      err = ch.Publish(
        "",
        q.Name,
        false,
        false,
        amqp.Publishing{
          DeliveryMode: amqp.Persistent,
          ContentType: "text/plain",
          Body: []byte(w),
        })

      erroMensage(err, "Falha ao publicar na fila")
		}

		fmt.Println(m, " - ")
		msec, _ := strconv.Atoi(s)
		time.Sleep(time.Duration(msec) * time.Millisecond)

		sender.Send("", 0)

	}

}
