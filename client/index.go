package main

import (
	zmq "github.com/pebbe/zmq4"

	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var m = make(map[string]int)

func msgErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

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
			defer wg.Done()
			/* if err := lerArquivo(fileName); err != nil {
							fmt.Println(err.Error())
			      } */
			file, err := os.Open(fileName)
			msgErro(err, "Erro ao abrir arquivo")
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				palavra := strings.ToLower(scanner.Text())
				palavra = removeSpecialCharacters(palavra)
				sender.Send(palavra, 0)
			}

		}(f)
	}
	wg.Wait()
	time.Sleep(time.Second)
}

/* func lerArquivo(fileName string) error {
	wg.Add(1)

	file, err := os.Open(fileName)
	if err != nil {
		return errors.New("Erro ao ler arquivo")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
    palavra := strings.ToLower(scanner.Text())
    palavra = removeSpecialCharacters(palavra)
	}

	wg.Done()
	return nil

}
*/
func removeSpecialCharacters(s string) string {
	reg, err := regexp.Compile("[[:punct:]0-9]+")
	msgErro(err, "Erro ao Copila Regex")
	processedText := reg.ReplaceAllString(s, "")

	return processedText

}
