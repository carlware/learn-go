package main

import (
	"fmt"
	"time"
)

var petNames = []string{"filomena", "puma", "pantera", "robertina", "nala", "alfonsina"}

func main() {

	chanOwner := func() chan string {
		resultSream := make(chan string)

		go func() {
			defer close(resultSream)
			for _, name := range petNames {
				time.Sleep(100 * time.Millisecond)
				resultSream <- name
			}
		}()

		return resultSream
	}

	resultStream := chanOwner()
	for name := range resultStream {
		fmt.Println("Pet name: ", name)
	}
}
