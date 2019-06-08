package main

import (
	zmq "github.com/pebbe/zmq4"

	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {

	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Bind("tcp://*:5557")

	//  Socket to send start of batch message on
	sink, _ := zmq.NewSocket(zmq.PUSH)
	defer sink.Close()
	sink.Connect("tcp://localhost:5558")

	sink.Send("0", 0)

	for _, f := range os.Args[1:] {
		wg.Add(1)

		go func(fileName string) {
			if err := lerArquivo(fileName, sender); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()

}

func lerArquivo(fileName string, sender *zmq.Socket) error {
	wg.Add(1)
	file, err := os.Open(fileName)
	if err != nil {
		return errors.New("Erro ao ler arquvo")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.)

	for scanner.Scan() {
		palavra := strings.ToLower(scanner.Text())
		if palavra != "" {
			sender.Send(palavra, 0)
		}

	}

	wg.Done()
	return nil

}
