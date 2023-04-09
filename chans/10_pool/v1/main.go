package main

import (
	"fmt"
	"sync"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new instance")
			return struct {}{}
		},
	}

	p.Get()
	instance := p.Get()
	p.Put(instance)
	p.Get()
	p.Put(instance)
	p.Get()
	p.Put(instance)
	p.Get()
	p.Put(instance)
	p.Get()
	p.Put(instance)
	p.Get()
}
