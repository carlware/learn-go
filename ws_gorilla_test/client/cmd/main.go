package main

import (
	"context"
	"fmt"
	"github.com/carlware/ws/client"
	"log"
	"sync"
)

const (
	host = "ws://localhost:8080/ws"
	hostTicker = "wss://arbitrage.ambitionsystems.com/ws/tickers"
	hostOpp = "wss://arbitrage.ambitionsystems.com/ws/opportunities"
)


func main() {
	//ctx := context.Background()
	//c, err := client.NewClient(hostOpp)
	//if err != nil {
	//	log.Printf("failed to create new client: %s", err)
	//}
	//
	//ch, lErr := c.Listen(ctx)
	//if lErr != nil {
	//	log.Printf("failed to lister: %s", lErr)
	//}
	//
	//for msg := range ch {
	//	fmt.Printf("message recv: %s\n", msg)
	//}
	fmt.Println("main")
	testWss(hostOpp)
}

func testWss(host string) {
	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(220)
	start := make(chan struct{})
	for i := 0; i < 220; i++ {
		go func(wg *sync.WaitGroup, id int) {
			c, err := client.NewClient(host)
			if err != nil {
				log.Printf("failed to create new client: %s", err)
			}

			<-start
			ch, lErr := c.Listen(ctx)
			if lErr != nil {
				log.Printf("failed to lister: %s", lErr)
			}

			for msg := range ch {
				_ = msg
				fmt.Printf("message recv from goroutine %d\n", id)
			}
			wg.Done()
		}(&wg, i)
	}

	close(start)
	wg.Wait()
}
