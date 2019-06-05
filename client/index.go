package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for _, f := range os.Args[1:] {
		wg.Add(1)

		go func(fileName string) {
			if err := lerArquivo(fileName); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
}

func lerArquivo(fileName string) error {
	wg.Add(1)
	file, err := os.Open(fileName)
	if err != nil {
		return errors.New("Erro ao ler arquivo")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		palavra := strings.ToLower(scanner.Text())
		imprimir(palavra)

	}

	wg.Done()
	return nil

	/*   func dividirarquivo (file string) error {
	     } */

}

func removeSpecialCharacters(s string){
	reg, err := regexp.Compile("[[:punct:]0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedText := reg.ReplaceAllString(s, "")

	fmt.Printf(processedText)

}


func imprimir(r string) {
	removeSpecialCharacters(r)
	//fmt.Println(r+ "gdusidh")
}

