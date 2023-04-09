package main

import (
	"fmt"
)

func main() {

	E1()
}

func E1() {
	done := make(chan struct{})
	stringStream := make(chan string)

	for _, s := range []string{"a", "b", "c"} {
		fmt.Println("looping", s)
		select {
		case <-done:
			fmt.Println("done")
			return
		case stringStream <- s:
			fmt.Println("pushed new val to chan")
		}
	}
}
