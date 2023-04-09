package main

import "fmt"

func main() {
	lucky := 0
	for i := 0; i < 1000000; i++ {
		first := i % 1000
		second := i / 1000

		if sumDigits(first) == sumDigits(second) {
			lucky += 1
		}
	}
	fmt.Println("lucky tickets", lucky)
}

func sumDigits(n int) int {
	u := n % 10
	d := ((n - u) % 100) / 10
	c := ((n - d - u) % 1000) / 100
	return u + d + c
}
