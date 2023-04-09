package main

import (
	"fmt"
	"math/rand"
	"time"
)

var take = func(
	done <-chan interface{},
	valueStream <-chan interface{},
	num int,
) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

var repeatFn = func(
	done <-chan interface{},
	fn func() interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

var toInt = func(done <-chan interface{}, ints <-chan interface{}) <-chan int {
	intChan := make(chan int)
	go func() {
		defer close(intChan)
		for {
			select {
			case <-done:
				return
			case val, _ := <-ints:
				intVal, ok := val.(int)
				if ok {
					intChan <- intVal
				}
			}
		}
	}()
	return intChan
}

var primeFinder = func(
	done <-chan interface{},
	ints <-chan int,
) <-chan interface{} {
	chanInts := make(chan interface{})
	go func() {
		defer close(chanInts)
	loop:
		for {
			select {
			case <-done:
				return
			case val, ok := <-ints:
				if !ok || val == 1 {
					continue loop
				}
				for i := 2; i < val; i++ {
					if val%i == 0 {
						continue loop
					}
				}
				chanInts <- val
			}
		}
	}()

	return chanInts
}

func main() {
	rand := func() interface{} { return rand.Intn(500000000) }
	done := make(chan interface{})
	defer close(done)
	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes:")

	for prime := range take(done, primeFinder(done, randIntStream), 5) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
