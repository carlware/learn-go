package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hola")
	}()

	time.Sleep(1 * time.Second)
}
