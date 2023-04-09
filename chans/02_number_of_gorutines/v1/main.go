package main

import (
	"fmt"
	"runtime"
)

func main() {
	gN := runtime.NumGoroutine()

	fmt.Println("gorutines:", gN)
}
