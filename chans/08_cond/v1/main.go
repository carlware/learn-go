package main

import (
	"fmt"
	"sync"
	"time"
)

var removeFromQueue = func(d time.Duration, cond *sync.Cond, q *[]interface{}) {
	time.Sleep(d)
	fmt.Println("before lock")
	cond.L.Lock()
	*q = (*q)[1:]
	fmt.Println("element removed from queue")
	cond.L.Unlock()
	fmt.Println("signaling ")
	cond.Signal()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

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
		go removeFromQueue(1*time.Second, cond, &queue)
		cond.L.Unlock()
	}
}