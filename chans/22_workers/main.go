package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	dateLayout = "15:04:05"
)

func main() {
	cxt, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2200 * time.Millisecond)
		fmt.Printf("Canceling %s\n", time.Now().Format("15:04:05"))
		cancel()
	}()

	wg := sync.WaitGroup{}

	strChan := t1(cxt)

	wg.Add(3)
	t2(cxt, strChan, &wg)

	wg.Wait()
	fmt.Println("finish")
}

func t1(ctx context.Context) <-chan string {
	strChan := make(chan string)
	go func() {
		for _, m := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"} {
			select {
			case <-ctx.Done():
				break
			default:
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				fmt.Printf("Sending %s message %s\n", time.Now().Format(dateLayout), m)
				strChan <- m
			}
		}
		close(strChan)
	}()
	return strChan
}

func t2(ctx context.Context, strChan <-chan string, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			for m := range strChan {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Printf("Proceesing %s message %s\n", time.Now().Format(dateLayout), m)
					t3(ctx, m)
					//time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
					//fmt.Printf("message %s\n", m)
				}
			}
			fmt.Printf("finish worker %d\n", id)
		}(i)
	}
}

func t3(ctx context.Context, message string) {
	var counter int
	ticker := time.NewTicker(300 * time.Millisecond)
	for range ticker.C {
		select {
		case <-ctx.Done():
			return
		default:
			if counter == 3 {
				fmt.Printf("Printing %s message %s\n", time.Now().Format("15:04:05"), message)
				return
			}
			ticker.Reset(time.Duration(rand.Intn(300)+1) * time.Millisecond)
			counter++
		}
	}
}
