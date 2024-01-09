package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя файла: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}

	reader = bufio.NewReader(file)

}
