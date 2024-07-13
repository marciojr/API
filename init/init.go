package main

import (
	"fmt"

	"github.com/message"
)

func main() {
	msg, _ := message.ShowMessage("Golang")
	fmt.Println(msg)

	_, error := message.ShowMessage("")
	fmt.Println(error)
}
