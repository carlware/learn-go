package main

import (
	"fmt"
	"math"
	"os"
	"runtime/pprof"
)

// fib1 calculates the Fibonacci series using recursion
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

// fib2 calculates the Fibonacci series using a golden ratio
func fib2(n int) int {
	phi := (1 + math.Sqrt(5)) / 2
	return int(math.Round(math.Pow(phi, float64(n)) / math.Sqrt(5)))
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	n := 50
	fmt.Printf("Fibonacci number at index %d using fib1: %d\n", n, fib1(n))
	fmt.Printf("Fibonacci number at index %d using fib2: %d\n", n, fib2(n))
}
