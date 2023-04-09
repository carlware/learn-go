package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(d time.Duration) {
		time.Sleep(d)
		cond.L.Lock()
		queue = queue[1:]
		fmt.Println("element removed from queue")
		cond.L.Unlock()
		cond.Signal()
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		fmt.Println("len", len(queue))
		for len(queue) == 2 {
			fmt.Println("waiting...")
			cond.Wait()
			fmt.Println("stop waiting")
		}
		fmt.Println("adding new element to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second)
		cond.L.Unlock()
	}
}