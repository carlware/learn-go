package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Println("begun", i)
		}(i)
	}

	fmt.Println("start")
	// don't <- struct{}{}
	close(begin)
	wg.Wait()
}
