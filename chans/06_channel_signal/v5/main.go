package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const MaxGorutines = 9
const maxSleep = 10

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("starting monitor")
	go monitor(ctx)

	for i := 0; i < MaxGorutines; i++ {
		go task(ctx, i)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	s := <-sig
	fmt.Println("Got signal:", s)

	cancel()
	fmt.Println("Grace period")
	time.Sleep(2 * time.Second)
	fmt.Println("Shutdown")
}

func task(ctx context.Context, taskId int) {
	timeToSleep := rand.Int()%maxSleep + 1
	timer := time.NewTimer(time.Duration(timeToSleep) * time.Second)
	fmt.Printf("Start task: %d sleep: time %d at %s \n", taskId, timeToSleep, time.Now().Format("15:04:05"))

	go func() {
		for {
			select {
			case time := <-timer.C:
				fmt.Printf("finish task %d at %s \n", taskId, time.Format("15:04:05"))
				return
			case <-ctx.Done():
				fmt.Printf("force to shutdown task %d at %s \n", taskId, time.Now().Format("15:04:05"))
				return
			}
		}
	}()
}

func monitor(ctx context.Context) {
	tick := time.Tick(1 * time.Second)
outer:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopping monitor")
			break outer
		case <-tick:
			gN := runtime.NumGoroutine()
			fmt.Println("gorutines:", gN)
		}
	}
}
