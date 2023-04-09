package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	var counter sync.WaitGroup

	for i := 0; i < 5; i++ {
		counter.Add(1)
		go func() {
			defer counter.Done()
			count++
		}()
	}

	counter.Wait()
	fmt.Printf("Counter: %d \n", count)
}
