package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{ Clicked: sync.NewCond(&sync.Mutex{}) }

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		fmt.Println("register new subscription")
		go func() {
			goroutineRunning.Done()
			fmt.Println("locking event")
			c.L.Lock()
			defer c.L.Unlock()
			fmt.Println("signal event")
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
		fmt.Println("stop subscription")
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
