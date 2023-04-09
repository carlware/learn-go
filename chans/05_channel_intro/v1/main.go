package main

import (
	"fmt"
	"time"
)

func main() {
	stringStream := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		stringStream <- "Hello world"
	}()

	fmt.Println("Waiting for the message")
	fmt.Printf("Stream message: %s \n", <-stringStream)
}
