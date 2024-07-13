package main

import (
	"fmt"
	"log"

	"github.com/message"
)

func main() {

	log.SetPrefix("init: ")
	log.SetFlags(0)

	msg, err := message.ShowMessage("Golang is a great lang")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
