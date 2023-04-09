package main

import (
	"fmt"
	"time"
)

func main() {
	stringStream := make(chan string)
	close(stringStream)

	go func() {
		defer close(stringStream)
		time.Sleep(1 * time.Second)
		stringStream <- "Hello world"
	}()

	fmt.Println("Waiting for the message")

	message, ok := <-stringStream
	fmt.Printf("Stream message: %s %v \n", message, ok)
}
