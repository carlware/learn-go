package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	intStream := make(chan int)
	limitChan := make(chan struct{}, 1)

	s := struct {
		sync.Mutex
		closed bool
	}{}

	var wg sync.WaitGroup
	wg.Add(2)

	go func(w *sync.WaitGroup) {
		defer func() {
			w.Done()
			fmt.Println( "Producer Done.")
		}()
		for i := 0; i < 10; i++ {
			fmt.Printf( "Sending: %d\n", i)

			go func(ii int) {
				s.Lock()
				defer s.Lock()
				if !s.closed {
					intStream <- ii
					fmt.Println("sent", ii)
				}
			}(i)
		}
		time.Sleep(3 * time.Second)
		for i := 0; i < 10; i++ {
			fmt.Printf( "Sending: %d\n", i)

			go func(ii int) {
				s.Lock()
				defer s.Lock()
				if !s.closed {
					intStream <- ii
					fmt.Println("sent", ii)
				}
			}(i)
		}
	}(&wg)

	ctx := context.Background()
	go func(w *sync.WaitGroup) {
		defer w.Done()
		for  {
			select {
			case <-ctx.Done():
				return
			case integer, ok := <-intStream:
				fmt.Printf( "Received %v.%v\n", integer, ok)
				if !ok {
					return
				}

				go func(i int) {
					defer func() {
						<-limitChan
					}()
					fmt.Println()
					fmt.Println("doing some work", i)
					time.Sleep(1 * time.Second)
					fmt.Println("finished some work", i)
				}(integer)
				limitChan <- struct{}{}
			default:

			}
		}
	}(&wg)

	go func() {
		t := time.NewTimer(1 * time.Second)
		<-t.C
		limitChan = make(chan struct{})
		intStream = make(chan int)
	}()

	wg.Wait()
}
