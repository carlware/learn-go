package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

func increment() {
	lock.Lock()
	defer lock.Unlock()
	count++
}

func decrement() {
	lock.Lock()
	defer lock.Unlock()
	count--
}

func main() {
	var counter sync.WaitGroup

	for i := 0; i < 5; i++ {
		counter.Add(1)
		go func() {
			defer counter.Done()
			increment()
		}()
	}

	for i := 0; i < 5; i++ {
		counter.Add(1)
		go func() {
			defer counter.Done()
			decrement()
		}()
	}

	counter.Wait()
	fmt.Printf("Counter: %d \n", count)
}
