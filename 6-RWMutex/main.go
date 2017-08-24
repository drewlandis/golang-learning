package main

import (
	"fmt"
	"sync"
)

// These mutex locks are really to prevent cross-go-routines from changing the value...
// So this example below doesn't really make sense
type hits struct {
	sync.RWMutex
	n int
}

func main() {
	h := hits{n: 1}
	fmt.Println("h.n:", h.n)
	h.Lock()
	h.n++
	h.Unlock()
	fmt.Println("h.n:", h.n)
	h.n++
	fmt.Println("h.n:", h.n)
}
