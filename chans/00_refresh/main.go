package main

import "fmt"

func bar() {
	fmt.Println("bar")
}

func main() {
	var name string
	var foo int

	petName := "patito"

	bar()
	fmt.Println("name", petName)
	fmt.Println("zero values", name, foo)

	func() {
		fmt.Println("closure")
	}()
}
