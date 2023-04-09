package main

import (
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

const MaxGorutines = 10
const maxSleep = 10

func main() {
	for i := 0; i < MaxGorutines; i++ {
		go func() {
			s := rand.Int() % maxSleep
			time.Sleep(time.Duration(s) * time.Second)
		}()
	}

	go monitor()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	s := <-sig
	fmt.Println("Got signal:", s)

	fmt.Println("Shutdown")
}

func monitor() {
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-tick:
			gN := runtime.NumGoroutine()
			fmt.Println("gorutines:", gN)
		}
	}
}
