package main

import "github.com/marciojr/API/go_gin"

func main() {
	println("Starting our financial manager web service...")
	go_gin.Initialize()
}
